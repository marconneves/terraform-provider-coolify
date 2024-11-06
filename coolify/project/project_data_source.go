package project

import (
	"context"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &ProjectDataSource{}

func NewProjectDataSource() datasource.DataSource {
	return &ProjectDataSource{}
}

type ProjectDataSource struct {
	client *coolify_sdk.Sdk
}

func (d *ProjectDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (d *ProjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Project by ID or Name data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Project identifier",
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Project name",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Project description",
				Computed:            true,
			},
			"environments": schema.ListNestedAttribute{
				MarkdownDescription: "List of environments associated with the project",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							MarkdownDescription: "Environment identifier",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Environment name",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Environment description",
							Computed:            true,
						},
						"project_id": schema.Int64Attribute{
							MarkdownDescription: "Associated project identifier",
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
				},
			},
		},
	}
}

func (d *ProjectDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	resp.Diagnostics.Append(configure.ConfigureClient(ctx, req, &d.client)...)
}

func (d *ProjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var project ProjectModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &project)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectSaved, diags := readProject(ctx, *d.client, project.Id, project.Name)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mapProjectDataSourceModel(&project, projectSaved)

	tflog.Trace(ctx, "Successfully read team data", map[string]interface{}{
		"project_id": projectSaved.ID,
		"name":       projectSaved.Name,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &project)...)
}
