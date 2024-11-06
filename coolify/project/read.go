package project

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func readProject(ctx context.Context, client coolify_sdk.Sdk, id types.String, name types.String) (*coolify_sdk.Project, diag.Diagnostics) {
	var diags diag.Diagnostics

	var project *coolify_sdk.Project
	var err error

	if !id.IsNull() {
		project, err = fetchProjectByID(client, ctx, id.ValueString())
	} else if !name.IsNull() {
		project, err = fetchProjectByName(client, ctx, name.ValueString())
	} else {
		diags.AddError("Configuration Error", "Either 'id' or 'name' must be specified.")
		return nil, diags
	}

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read project data: %s", err))
		return nil, diags
	}

	if project == nil {
		diags.AddError("Not Found", "No project found with the given ID or name")
		return nil, diags
	}

	return project, diags
}
