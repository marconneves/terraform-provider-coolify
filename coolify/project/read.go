package project

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func readProject(ctx context.Context, client coolify_sdk.Sdk, id types.String, name *types.String) (*coolify_sdk.Project, diag.Diagnostics) {
	var diags diag.Diagnostics

	var project *coolify_sdk.Project
	var err error

	if !id.IsNull() {
		project, err = fetchProjectByID(client, ctx, id.ValueString())
	} else if !name.IsNull() {
		project, err = fetchProjectByName(client, ctx, name.ValueString())
	} else {
		diags.AddError("Configuration Error", "Either 'id' or 'name' must be specified.")
		return nil, diags
	}

	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read project data: %s", err))
		return nil, diags
	}

	if project == nil {
		diags.AddError("Not Found", "No project found with the given ID or name")
		return nil, diags
	}

	return project, diags
}

func (d *ProjectDataSource) ReadProjectDatasource(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var project ProjectDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &project)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectSaved, diags := readProject(ctx, *d.client, project.Id, &project.Name)
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

func (r *ProjectResource) ReadProjectResource(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var project ProjectModel

	resp.Diagnostics.Append(req.State.Get(ctx, &project)...)
	if resp.Diagnostics.HasError() {
		return
	}

	projectSaved, diags := readProject(ctx, *r.client, project.Id, nil)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mapProjectResourceModel(&project, projectSaved)

	tflog.Trace(ctx, "Successfully read team data", map[string]interface{}{
		"project_id": projectSaved.ID,
		"name":       projectSaved.Name,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &project)...)
}
