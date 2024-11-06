package project_environment

import (
	"context"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &EnvironmentDataSource{}

func NewEnvironmentDataSource() datasource.DataSource {
	return &EnvironmentDataSource{}
}

type EnvironmentDataSource struct {
	client *coolify_sdk.Sdk
}

type EnvironmentDataSourceModel struct {
	Id          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
	ProjectUUID types.String `tfsdk:"project_uuid"`
	ProjectID   types.Int64  `tfsdk:"project_id"`
}

func (d *EnvironmentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_environment"
}

func (d *EnvironmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Environment data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "Environment identifier",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Environment name",
				Required:            true,
			},
			"project_uuid": schema.StringAttribute{
				MarkdownDescription: "Project unique identifier",
				Required:            true,
			},
			"project_id": schema.Int64Attribute{
				MarkdownDescription: "Project identifier",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Environment description",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "Environment creation timestamp",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "Environment last update timestamp",
				Computed:            true,
			},
		},
	}
}

func (d *EnvironmentDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	resp.Diagnostics.Append(configure.ConfigureClient(ctx, req, &d.client)...)
}

func (d *EnvironmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var environment EnvironmentModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &environment)...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentSaved, diags := readEnvironment(*d.client, environment.ProjectUUID, environment.Name)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mapEnvironmentModel(&environment, environmentSaved)

	tflog.Trace(ctx, "Successfully read team data", map[string]interface{}{
		"environment_id": environmentSaved.Id,
		"name":           environmentSaved.Name,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &environment)...)
}
