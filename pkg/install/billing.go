package install

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
)

func (i *Installer) ensureBillingRecord(ctx context.Context) error {
	return i.billing.Ensure(ctx, i.doc)
}
