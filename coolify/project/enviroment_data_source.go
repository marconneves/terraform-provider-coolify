package project

import (
	"context"
	"fmt"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"

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
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*coolify_sdk.Sdk)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *coolify_sdk.Sdk, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *EnvironmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data EnvironmentDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	environment, err := d.client.Project.Environment(data.ProjectUUID.ValueString(), data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Client Error",
			fmt.Sprintf("Unable to read environment data, got error: %s", err),
		)
		return
	}

	data.Id = types.Int64Value(int64(environment.Id))
	data.Name = types.StringValue(environment.Name)
	data.Description = types.StringValue(environment.Description)
	data.ProjectID = types.Int64Value(int64(environment.ProjectID))
	data.CreatedAt = types.StringValue(environment.CreatedAt.String())
	data.UpdatedAt = types.StringValue(environment.UpdatedAt.String())

	tflog.Trace(ctx, "read an environment data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
