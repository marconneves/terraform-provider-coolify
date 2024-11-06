package team_members

import (
	"context"
	"fmt"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &TeamMembersDataSource{}

func NewTeamMembersDataSource() datasource.DataSource {
	return &TeamMembersDataSource{}
}

type TeamMembersDataSource struct {
	client *coolify_sdk.Sdk
}

type TeamMembersDataSourceModel struct {
	TeamId  types.Int64   `tfsdk:"team_id"`
	Members []MemberModel `tfsdk:"members"`
}

type MemberModel struct {
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Email types.String `tfsdk:"email"`
}

func (d *TeamMembersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team_members"
}

func (d *TeamMembersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

func (d *TeamMembersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*coolify_sdk.Sdk)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *TeamMembersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data TeamMembersDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Convert int64 to int
	teamID := int(data.TeamId.ValueInt64())

	members, err := d.client.Team.Members(teamID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read team members, got error: %s", err),
		)
		return
	}

	for _, member := range *members {
		data.Members = append(data.Members, MemberModel{
			Id:    types.StringValue(fmt.Sprintf("%d", member.Id)),
			Name:  types.StringValue(member.Name),
			Email: types.StringValue(member.Email),
		})
	}

	tflog.Trace(ctx, "read team members data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
