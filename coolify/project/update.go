package project

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func (r *ProjectResource) UpdateProject(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProjectModel
	var state ProjectModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.client == nil {
		resp.Diagnostics.AddError("Client Error", "Client is not configured. Please ensure the provider is properly configured.")
		return
	}

	if state.Id.IsNull() || state.Id.ValueString() == "" {
		resp.Diagnostics.AddError("ID Missing", "UUID is required to update the project.")
		return
	}

	updateDTO := coolify_sdk.UpdateProjectDTO{
		Name:        plan.Name.ValueStringPointer(),
		Description: plan.Description.ValueStringPointer(),
	}

	err := r.client.Project.Update(state.Id.ValueString(), &updateDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update project, got error: %s", err))
		return
	}

	project, err := r.client.Project.Get(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read project after update, got error: %s", err))
		return
	}

	var newState ProjectModel
	newState.Id = types.StringValue(project.UUID)
	newState.Name = types.StringValue(project.Name)
	if project.Description != nil {
		newState.Description = types.StringValue(*project.Description)
	} else {
		newState.Description = types.StringNull()
	}

	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}
