package private_key

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func readPrivateKey(ctx context.Context, client coolify_sdk.Sdk, id types.String) (*coolify_sdk.PrivateKey, diag.Diagnostics) {
	var diags diag.Diagnostics

	var privateKey *coolify_sdk.PrivateKey
	var err error

	if !id.IsNull() {
		privateKey, err = fetchPrivateKeyByUUID(client, ctx, id.ValueString())
	} else {
		diags.AddError("Configuration Error", "Either 'id' or 'name' must be specified.")
		return nil, diags
	}

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read private key data: %s", err))
		return nil, diags
	}

	if privateKey == nil {
		diags.AddError("Not Found", "No private key found with the given ID or name")
		return nil, diags
	}

	return privateKey, diags
}
