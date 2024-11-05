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

var _ datasource.DataSource = &ProjectDataSource{}

func NewProjectDataSource() datasource.DataSource {
	return &ProjectDataSource{}
}

type ProjectDataSource struct {
	client *coolify_sdk.Sdk
}

type ProjectDataSourceModel struct {
	Id           types.String       `tfsdk:"id"`
	Name         types.String       `tfsdk:"name"`
	Description  types.String       `tfsdk:"description"`
	Environments []EnvironmentModel `tfsdk:"environments"`
}

type EnvironmentModel struct {
	Id          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	ProjectId   types.Int64  `tfsdk:"project_id"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
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

func (d *ProjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ProjectDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var project *coolify_sdk.Project
	var err error

	if !data.Id.IsNull() {
		project, err = d.client.Project.Get(fmt.Sprintf("%v", data.Id.ValueString()))
		if err != nil {
			resp.Diagnostics.AddError(
				"Client Error",
				fmt.Sprintf("Unable to read project data by ID, got error: %s", err),
			)
			return
		}
	} else if !data.Name.IsNull() {
		projects, err := d.client.Project.List()
		if err != nil {
			resp.Diagnostics.AddError(
				"Client Error",
				fmt.Sprintf("Unable to list projects, got error: %s", err),
			)
			return
		}

		for _, p := range *projects {
			if p.Name == data.Name.ValueString() {
				project = &p
				break
			}
		}

		if project == nil {
			resp.Diagnostics.AddError(
				"Not Found",
				fmt.Sprintf("No project found with name: %s", data.Name.ValueString()),
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

	data.Id = types.StringValue(project.UUID)
	data.Name = types.StringValue(project.Name)
	if project.Description != nil {
		data.Description = types.StringValue(*project.Description)
	}

	if project.Environments != nil {
		environments := make([]EnvironmentModel, len(project.Environments))
		for i, env := range project.Environments {
			var description types.String
			if env.Description != nil {
				description = types.StringValue(*env.Description)
			} else {
				description = types.StringNull()
			}

			environments[i] = EnvironmentModel{
				Id:          types.Int64Value(env.ID),
				Name:        types.StringValue(env.Name),
				Description: description,
				ProjectId:   types.Int64Value(env.ProjectId),
				CreatedAt:   types.StringValue(env.CreatedAt.String()),
				UpdatedAt:   types.StringValue(env.UpdatedAt.String()),
			}
		}
		data.Environments = environments
	}

	tflog.Trace(ctx, "read a project by ID or Name data source")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
