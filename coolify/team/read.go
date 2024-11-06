package team

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func readTeam(ctx context.Context, client coolify_sdk.Sdk, id types.Int64, name types.String) (*coolify_sdk.Team, diag.Diagnostics) {
	var diags diag.Diagnostics

	var team *coolify_sdk.Team
	var err error

	if !id.IsNull() {
		team, err = fetchTeamByID(client, ctx, id.ValueInt64())
	} else if !name.IsNull() {
		team, err = fetchTeamByName(client, ctx, name.ValueString())
	} else {
		diags.AddError("Configuration Error", "Either 'id' or 'name' must be specified.")
		return nil, diags
	}

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read team data: %s", err))
		return nil, diags
	}

	if team == nil {
		diags.AddError("Not Found", "No team found with the given ID or name")
		return nil, diags
	}

	return team, diags
}
