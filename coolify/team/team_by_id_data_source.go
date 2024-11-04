package team

import (
	"context"
	"fmt"

	"github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &TeamByIDDataSource{}

func NewTeamByIDDataSource() datasource.DataSource {
	return &TeamByIDDataSource{}
}

type TeamByIDDataSource struct {
	client *coolify_sdk.Sdk
}

type TeamByIDDataSourceModel struct {
	Id          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func (d *TeamByIDDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team_by_id"
}

func (d *TeamByIDDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Team by ID data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "Team identifier",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Team name",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Team description",
				Computed:            true,
			},
		},
	}
}

func (d *TeamByIDDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *TeamByIDDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data TeamByIDDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	team, err := d.client.Team.Get(int(data.Id.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read team data by ID, got error: %s", err),
		)
		return
	}

	data.Name = types.StringValue(team.Name)
	if team.Description != nil {
		data.Description = types.StringValue(*team.Description)
	}

	tflog.Trace(ctx, "read a team by ID data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
