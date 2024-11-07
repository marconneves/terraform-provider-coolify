// create.go
package project

import (
	"context"
	"fmt"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (r *ProjectResource) CreateProject(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ProjectModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Planned data before creation", map[string]interface{}{
		"name":        data.Name.ValueString(),
		"description": data.Description.ValueString(),
	})

	createDTO := coolify_sdk.CreateProjectDTO{
		Name:        data.Name.ValueStringPointer(),
		Description: data.Description.ValueStringPointer(),
	}

	projectID, err := r.client.Project.Create(&createDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create project, got error: %s", err))
		return
	}

	data.Id = types.StringValue(*projectID)
	tflog.Trace(ctx, "Created a project", map[string]interface{}{"project_id": projectID})

	tflog.Debug(ctx, "Data after project creation", map[string]interface{}{
		"id":          data.Id.ValueString(),
		"name":        data.Name.ValueString(),
		"description": data.Description.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Debug(ctx, "Project state saved to file after creation")
}
