package configure

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func ConfigureClient(ctx context.Context, req interface{}, client **coolify_sdk.Sdk) diag.Diagnostics {
	var diags diag.Diagnostics

	var providerData interface{}
	switch v := req.(type) {
	case datasource.ConfigureRequest:
		providerData = v.ProviderData
	case resource.ConfigureRequest:
		providerData = v.ProviderData
	default:
		diags.AddError("Unexpected Request Type", "Unsupported request type in ConfigureClient.")
		return diags
	}

	if providerData == nil {
		diags.AddWarning("Configuration Warning", "Provider data is nil. Skipping configuration.")
		return diags
	}

	c, ok := providerData.(*coolify_sdk.Sdk)
	if !ok {
		diags.AddError(
			"Unexpected Client Type",
			fmt.Sprintf("Expected *coolify_sdk.Sdk, got: %T. Please report this issue to the provider developers.", providerData),
		)
		return diags
	}

	*client = c
	return diags
}
