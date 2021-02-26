package dynamichelper

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/ugorji/go/codec"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"

	"github.com/Azure/ARO-RP/pkg/util/cmp"
)

type Interface interface {
	Refresh() error
	EnsureDeleted(ctx context.Context, groupKind, namespace, name string) error
	Ensure(ctx context.Context, objs ...*unstructured.Unstructured) error
}

type dynamicHelper struct {
	GVRResolver

	log *logrus.Entry

	restconfig *rest.Config
	dyn        dynamic.Interface
}

func New(log *logrus.Entry, restconfig *rest.Config) (Interface, error) {
	dh := &dynamicHelper{
		log:        log,
		restconfig: restconfig,
	}

	var err error

	dh.GVRResolver, err = NewGVRResolver(log, restconfig)
	if err != nil {
		return nil, err
	}

	dh.dyn, err = dynamic.NewForConfig(restconfig)
	if err != nil {
		return nil, err
	}

	return dh, nil
}

func (dh *dynamicHelper) EnsureDeleted(ctx context.Context, groupKind, namespace, name string) error {
	gvr, err := dh.Resolve(groupKind, "")
	if err != nil {
		return err
	}

	err = dh.dyn.Resource(*gvr).Namespace(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if errors.IsNotFound(err) {
		err = nil
	}
	return err
}

// Ensure is called by the operator deploy tool and individual controllers.  It
// is intended to ensure that an object matches a desired state.  It is tolerant
// of unspecified fields in the desired state (e.g. it will leave typically
// leave .status untouched).
func (dh *dynamicHelper) Ensure(ctx context.Context, objs ...*unstructured.Unstructured) error {
	for _, o := range objs {
		err := dh.ensureOne(ctx, o)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dh *dynamicHelper) ensureOne(ctx context.Context, o *unstructured.Unstructured) error {
	gvr, err := dh.Resolve(o.GroupVersionKind().GroupKind().String(), o.GroupVersionKind().Version)
	if err != nil {
		return err
	}

	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		existing, err := dh.dyn.Resource(*gvr).Namespace(o.GetNamespace()).Get(ctx, o.GetName(), metav1.GetOptions{})
		if errors.IsNotFound(err) {
			dh.log.Printf("Create %s", keyFuncO(o))
			_, err = dh.dyn.Resource(*gvr).Namespace(o.GetNamespace()).Create(ctx, o, metav1.CreateOptions{})
			return err
		}
		if err != nil {
			return err
		}

		o, changed, diff, err := merge(existing, o)
		if err != nil || !changed {
			return err
		}

		dh.log.Printf("Update %s: %s", keyFuncO(o), diff)

		_, err = dh.dyn.Resource(*gvr).Namespace(o.GetNamespace()).Update(ctx, o, metav1.UpdateOptions{})
		return err
	})
}

func diff(existing, o *unstructured.Unstructured) string {
	if o.GroupVersionKind().GroupKind().String() == "Secret" { // Don't show a diff if kind is Secret
		return ""
	}

	return cmp.Diff(existing.Object, o.Object)
}

// merge merges delta onto base using ugorji/go/codec semantics.  It returns the
// newly merged object (the inputs are untouched) plus a flag indicating if a
// change took place and a printable diff as appropriate
func merge(base, delta *unstructured.Unstructured) (*unstructured.Unstructured, bool, string, error) {
	copy := base.DeepCopy()

	h := &codec.JsonHandle{
		MapKeyAsString: true,
	}

	var b []byte
	err := codec.NewEncoderBytes(&b, h).Encode(delta.Object)
	if err != nil {
		return nil, false, "", err
	}

	err = codec.NewDecoderBytes(b, h).Decode(&copy.Object)
	if err != nil {
		return nil, false, "", err
	}

	// all new objects have a null creationTimestamp that causes every object to
	// be updated.
	copy.SetCreationTimestamp(base.GetCreationTimestamp())

	status, found, err := unstructured.NestedMap(base.Object, "status")
	if err == nil && found {
		err = unstructured.SetNestedMap(copy.Object, status, "status")
		if err != nil {
			return nil, false, "", err
		}
	} else {
		// prevent empty status objects from causing problems
		unstructured.RemoveNestedField(copy.Object, "status")
	}

	return copy, !reflect.DeepEqual(base, copy), diff(base, copy), nil
}

func keyFuncO(o *unstructured.Unstructured) string {
	return keyFunc(o.GroupVersionKind().GroupKind(), o.GetNamespace(), o.GetName())
}

func keyFunc(gk schema.GroupKind, namespace, name string) string {
	s := gk.String()
	if namespace != "" {
		s += "/" + namespace
	}
	s += "/" + name

	return s
}
