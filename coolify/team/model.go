package team

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

type TeamModel struct {
	Id          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func mapTeamModel(team *TeamModel, teamSaved *coolify_sdk.Team) {
	team.Id = types.Int64Value(int64(teamSaved.Id))
	team.Name = types.StringValue(teamSaved.Name)
	if teamSaved.Description != nil {
		team.Description = types.StringValue(*teamSaved.Description)
	}
}
