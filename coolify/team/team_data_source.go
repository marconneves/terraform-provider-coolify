package team

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ datasource.DataSource = &TeamDataSource{}

func NewTeamDataSource() datasource.DataSource {
	return &TeamDataSource{}
}

type TeamDataSource struct {
	client *coolify_sdk.Sdk
}

func (d *TeamDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team"
}

func (d *TeamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Team by ID or Name data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "Team identifier",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Team name",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Team description",
				Computed:            true,
			},
		},
	}
}

func (d *TeamDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	resp.Diagnostics.Append(configure.ConfigureClient(ctx, req, &d.client)...)
}

func (d *TeamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var teamModel TeamModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &teamModel)...)
	if resp.Diagnostics.HasError() {
		return
	}

	team, diags := readTeam(ctx, *d.client, teamModel.Id, teamModel.Name)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mapTeamModel(&teamModel, team)

	tflog.Trace(ctx, "Successfully read team data", map[string]interface{}{
		"team_id": team.Id,
		"name":    team.Name,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &teamModel)...)
}
