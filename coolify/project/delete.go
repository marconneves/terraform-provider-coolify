package project

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (r *ProjectResource) DeleteProject(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ProjectModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Project.Delete(data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete project, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted project", map[string]interface{}{"project_id": data.Id.ValueString()})
}
