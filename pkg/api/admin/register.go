package admin

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/util/deployment"
)

// APIVersion contains a version string as it will be used by clients
const APIVersion = "admin"

func init() {
	api.APIs[APIVersion] = &api.Version{
		OpenShiftClusterConverter: func() api.OpenShiftClusterConverter {
			return &openShiftClusterConverter{}
		},
		OpenShiftClusterStaticValidator: func(string, string, deployment.Mode, string) api.OpenShiftClusterStaticValidator {
			return &openShiftClusterStaticValidator{}
		},
	}
}
