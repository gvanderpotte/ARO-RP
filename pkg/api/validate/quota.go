package validate

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/sirupsen/logrus"

	"github.com/Azure/ARO-RP/pkg/api"
	"github.com/Azure/ARO-RP/pkg/env"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/compute"
	"github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/network"
)

func addRequiredResources(requiredResources map[string]int, vmSize api.VMSize, count int) error {
	requiredResources["virtualMachines"] += count
	requiredResources["PremiumDiskCount"] += count
	switch vmSize {
	case api.VMSizeStandardD2sV3:
		requiredResources["standardDSv3Family"] += (count * 2)
		requiredResources["cores"] += (count * 2)

	case api.VMSizeStandardD4asV4:
		requiredResources["standardDASv4Family"] += (count * 4)
		requiredResources["cores"] += (count * 4)
	case api.VMSizeStandardD8asV4:
		requiredResources["standardDASv4Family"] += (count * 8)
		requiredResources["cores"] += (count * 8)
	case api.VMSizeStandardD16asV4:
		requiredResources["standardDASv4Family"] += (count * 16)
		requiredResources["cores"] += (count * 16)
	case api.VMSizeStandardD32asV4:
		requiredResources["standardDASv4Family"] += (count * 32)
		requiredResources["cores"] += (count * 32)

	case api.VMSizeStandardD4sV3:
		requiredResources["standardDSv3Family"] += (count * 4)
		requiredResources["cores"] += (count * 4)
	case api.VMSizeStandardD8sV3:
		requiredResources["standardDSv3Family"] += (count * 8)
		requiredResources["cores"] += (count * 8)
	case api.VMSizeStandardD16sV3:
		requiredResources["standardDSv3Family"] += (count * 16)
		requiredResources["cores"] += (count * 16)
	case api.VMSizeStandardD32sV3:
		requiredResources["standardDSv3Family"] += (count * 32)
		requiredResources["cores"] += (count * 32)

	case api.VMSizeStandardE4sV3:
		requiredResources["standardESv3Family"] += (count * 4)
		requiredResources["cores"] += (count * 4)
	case api.VMSizeStandardE8sV3:
		requiredResources["standardESv3Family"] += (count * 8)
		requiredResources["cores"] += (count * 8)
	case api.VMSizeStandardE16sV3:
		requiredResources["standardESv3Family"] += (count * 16)
		requiredResources["cores"] += (count * 16)
	case api.VMSizeStandardE32sV3:
		requiredResources["standardESv3Family"] += (count * 32)
		requiredResources["cores"] += (count * 32)

	case api.VMSizeStandardF4sV2:
		requiredResources["standardFSv2Family"] += (count * 4)
		requiredResources["cores"] += (count * 4)
	case api.VMSizeStandardF8sV2:
		requiredResources["standardFSv2Family"] += (count * 8)
		requiredResources["cores"] += (count * 8)
	case api.VMSizeStandardF16sV2:
		requiredResources["standardFSv2Family"] += (count * 16)
		requiredResources["cores"] += (count * 16)
	case api.VMSizeStandardF32sV2:
		requiredResources["standardFSv2Family"] += (count * 32)
		requiredResources["cores"] += (count * 32)

	default:
		//will only happen if pkg/api verification allows new VMSizes
		return fmt.Errorf("unexpected node VMSize %s", vmSize)
	}
	return nil
}

type AzureQuotaValidator interface {
	Validate(context.Context) error
}

type quotaValidator struct {
	log *logrus.Entry
	env env.Interface

	oc             *api.OpenShiftCluster
	spComputeUsage compute.UsageClient
	spNetworkUsage network.UsageClient
}

func NewAzureQuotaValidator(ctx context.Context, log *logrus.Entry, env env.Interface, oc *api.OpenShiftCluster, subscriptionDoc *api.SubscriptionDocument) (AzureQuotaValidator, error) {
	r, err := azure.ParseResourceID(oc.ID)
	if err != nil {
		return nil, err
	}

	spAuthorizer, err := validateServicePrincipalProfile(ctx, log, &oc.Properties.ServicePrincipalProfile, subscriptionDoc.Subscription.Properties.TenantID, env.Environment().ResourceManagerEndpoint)
	if err != nil {
		return nil, err
	}

	validator := &quotaValidator{
		log: log,
		env: env,

		oc:             oc,
		spComputeUsage: compute.NewUsageClient(env.Environment(), r.SubscriptionID, spAuthorizer),
		spNetworkUsage: network.NewUsageClient(env.Environment(), r.SubscriptionID, spAuthorizer),
	}

	return validator, nil
}

// Validate checks usage quotas vs. resources required by cluster before cluster
// creation
func (qv *quotaValidator) Validate(ctx context.Context) error {
	qv.log.Print("ValidateQuotas")

	requiredResources := map[string]int{}
	err := addRequiredResources(requiredResources, qv.oc.Properties.MasterProfile.VMSize, 3)
	if err != nil {
		return err
	}
	//worker node resource calculation
	for _, w := range qv.oc.Properties.WorkerProfiles {
		err = addRequiredResources(requiredResources, w.VMSize, w.Count)
		if err != nil {
			return err
		}
	}

	//Public IP Addresses minimum requirement: 2 for ARM template deployment and 1 for kube-controller-manager
	requiredResources["PublicIPAddresses"] = 3

	//check requirements vs. usage
	// we're only checking the limits returned by the Usage API and ignoring usage limits missing from the results
	// rationale:
	// 1. if the Usage API doesn't send a limit because a resource is no longer limited, RP will continue cluster creation without impact
	// 2. if the Usage API doesn't send a limit that is still enforced, cluster creation will fail on the backend and we will get an error in the RP logs
	computeUsages, err := qv.spComputeUsage.List(ctx, qv.oc.Location)
	if err != nil {
		return err
	}

	qv.log.Print("Number of compute usages for location ", qv.oc.Location, ": ", len(computeUsages))
	for _, computeUsage := range computeUsages {
		qv.log.Print("Compute usage ", *computeUsage.Name.Value, ": ", *computeUsage.CurrentValue, " / ", *computeUsage.Limit)
		required, present := requiredResources[*computeUsage.Name.Value]
		if present && int64(required) > (*computeUsage.Limit-int64(*computeUsage.CurrentValue)) {
			return api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeResourceQuotaExceeded, "", "Resource quota of %s exceeded. Maximum allowed: %d, Current in use: %d, Additional requested: %d.", *computeUsage.Name.Value, *computeUsage.Limit, *computeUsage.CurrentValue, required)
		}
	}

	netUsages, err := qv.spNetworkUsage.List(ctx, qv.oc.Location)
	if err != nil {
		return err
	}
	qv.log.Print("Number of network usages for location ", qv.oc.Location, ": ", len(netUsages))
	for _, netUsage := range netUsages {
		qv.log.Print("Network usage ", *netUsage.Name.Value, ": ", *netUsage.CurrentValue, " / ", *netUsage.Limit)
		required, present := requiredResources[*netUsage.Name.Value]
		if present && int64(required) > (*netUsage.Limit-int64(*netUsage.CurrentValue)) {
			return api.NewCloudError(http.StatusBadRequest, api.CloudErrorCodeResourceQuotaExceeded, "", "Resource quota of %s exceeded. Maximum allowed: %d, Current in use: %d, Additional requested: %d.", *netUsage.Name.Value, *netUsage.Limit, *netUsage.CurrentValue, required)
		}
	}

	return nil
}
