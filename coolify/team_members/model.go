package team_members

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

type TeamMembersModel struct {
	TeamId  types.Int64   `tfsdk:"team_id"`
	Members []MemberModel `tfsdk:"members"`
}

type MemberModel struct {
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Email types.String `tfsdk:"email"`
}

func mapTeamMembersModel(teamMembers *TeamMembersModel, members *[]coolify_sdk.Member) {

	for _, member := range *members {
		teamMembers.Members = append(teamMembers.Members, MemberModel{
			Id:    types.StringValue(fmt.Sprintf("%d", member.Id)),
			Name:  types.StringValue(member.Name),
			Email: types.StringValue(member.Email),
		})
	}

}
