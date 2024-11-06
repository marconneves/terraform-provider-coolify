package team_members

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func readTeamMembers(client coolify_sdk.Sdk, id types.Int64) (*[]coolify_sdk.Member, diag.Diagnostics) {
	var diags diag.Diagnostics

	var members *[]coolify_sdk.Member
	var err error

	if !id.IsNull() {
		teamID := int(id.ValueInt64())

		members, err = client.Team.Members(teamID)
	} else {
		diags.AddError("Configuration Error", "Either 'id' or 'name' must be specified.")
		return nil, diags
	}

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read team data: %s", err))
		return nil, diags
	}

	if members == nil {
		diags.AddError("Not Found", "No team found members")
		return nil, diags
	}

	return members, diags
}
