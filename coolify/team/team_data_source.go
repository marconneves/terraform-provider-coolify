package team

import (
	"context"
	"fmt"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &TeamDataSource{}

func NewTeamDataSource() datasource.DataSource {
	return &TeamDataSource{}
}

type TeamDataSource struct {
	client *coolify_sdk.Sdk
}

type TeamDataSourceModel struct {
	Id          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
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

func (d *TeamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data TeamDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var team *coolify_sdk.Team
	var err error

	if !data.Id.IsNull() {
		team, err = d.client.Team.Get(int(data.Id.ValueInt64()))
		if err != nil {
			resp.Diagnostics.AddError(
				"Client Error",
				fmt.Sprintf("Unable to read team data by ID, got error: %s", err),
			)
			return
		}
	} else if !data.Name.IsNull() {
		teams, err := d.client.Team.List()
		if err != nil {
			resp.Diagnostics.AddError(
				"Client Error",
				fmt.Sprintf("Unable to list teams, got error: %s", err),
			)
			return
		}

		for _, t := range *teams {
			if t.Name == data.Name.ValueString() {
				team = &t
				break
			}
		}

		if team == nil {
			resp.Diagnostics.AddError(
				"Not Found",
				fmt.Sprintf("No team found with name: %s", data.Name.ValueString()),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Configuration Error",
			"Either 'id' or 'name' must be specified.",
		)
		return
	}

	data.Id = types.Int64Value(int64(team.Id))
	data.Name = types.StringValue(team.Name)
	if team.Description != nil {
		data.Description = types.StringValue(*team.Description)
	}

	tflog.Trace(ctx, "read a team by ID or Name data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
