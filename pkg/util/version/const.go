package version

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"github.com/Azure/ARO-RP/pkg/api"
)

const InstallArchitectureVersion = api.ArchitectureVersionV2

var GitCommit = "unknown"

// InstallStream describes stream we are defaulting to for all new clusters
var InstallStream = &Stream{
	Version:  NewVersion(4, 5, 22),
	PullSpec: "quay.io/openshift-release-dev/ocp-release@sha256:adb5ef06c54ff75ca9033f222ac5e57f2fd82e49bdd84f737d460ae542c8af60",
}

// Streams describes list of streams we support for upgrades
var (
	Streams = []*Stream{
		InstallStream,
		{
			Version:  NewVersion(4, 4, 31),
			PullSpec: "quay.io/openshift-release-dev/ocp-release@sha256:679db43a28a42fc41784ea3d4976d9d60cd194757cfdbea6137d6d0093db8c8d",
		},
		{
			Version:  NewVersion(4, 3, 40),
			PullSpec: "quay.io/openshift-release-dev/ocp-release@sha256:59cc585be7b4ad069a18f6f1a3309391e172192744ee65fa6e499c8b337edda4",
		},
	}
)

// FluentbitImage contains the location of the Fluentbit container image
func FluentbitImage(acrDomain string) string {
	return acrDomain + "/fluentbit:1.3.9-1"
}

// MdmImage contains the location of the MDM container image
func MdmImage(acrDomain string) string {
	return acrDomain + "/genevamdm:master_52"
}

// MdsdImage contains the location of the MDSD container image
func MdsdImage(acrDomain string) string {
	return acrDomain + "/genevamdsd:master_20201121.1"
}

func RouteFixImage(acrDomain string) string {
	return acrDomain + "/routefix:c5c4a5db"
}
