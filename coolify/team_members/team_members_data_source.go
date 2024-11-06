package team_members

import (
	"context"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &TeamMembersDataSource{}

func NewTeamMembersDataSource() datasource.DataSource {
	return &TeamMembersDataSource{}
}

type TeamMembersDataSource struct {
	client *coolify_sdk.Sdk
}

func (t *TeamMembersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team_members"
}

func (t *TeamMembersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Team members data source",

		Attributes: map[string]schema.Attribute{
			"team_id": schema.Int64Attribute{
				MarkdownDescription: "Team identifier",
				Required:            true,
			},
			"members": schema.ListNestedAttribute{
				MarkdownDescription: "List of team members",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Member identifier",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Member name",
							Computed:            true,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: "Member email",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (t *TeamMembersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	resp.Diagnostics.Append(configure.ConfigureClient(ctx, req, &t.client)...)
}

func (t *TeamMembersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var teamMembers TeamMembersModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &teamMembers)...)

	if resp.Diagnostics.HasError() {
		return
	}

	members, diags := readTeamMembers(*t.client, teamMembers.TeamId)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mapTeamMembersModel(&teamMembers, members)

	tflog.Trace(ctx, "Successfully read team data", map[string]interface{}{
		"team_id": teamMembers.TeamId,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &teamMembers)...)
}
