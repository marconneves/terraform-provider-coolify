package server

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func readServer(ctx context.Context, client coolify_sdk.Sdk, id types.String, name types.String) (*coolify_sdk.Server, diag.Diagnostics) {
	var diags diag.Diagnostics

	var server *coolify_sdk.Server
	var err error

	if !id.IsNull() {
		server, err = fetchServerByID(client, ctx, id.ValueString())
	} else if !name.IsNull() {
		server, err = fetchServerByName(client, ctx, name.ValueString())
	} else {
		diags.AddError("Configuration Error", "Either 'id' or 'name' must be specified.")
		return nil, diags
	}

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read server data: %s", err))
		return nil, diags
	}

	if server == nil {
		diags.AddError("Not Found", "No server found with the given ID or name")
		return nil, diags
	}

	return server, diags
}
