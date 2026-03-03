package server

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func (r *ServerResource) UpdateServer(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ServerModel
	var state ServerModel

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

	if state.ID.IsNull() || state.ID.ValueString() == "" {
		resp.Diagnostics.AddError("ID Missing", "UUID is required to update the server.")
		return
	}

	updateDTO := coolify_sdk.UpdateServerDTO{}
	if !plan.Name.Equal(state.Name) {
		updateDTO.Name = plan.Name.ValueString()
	}
	if !plan.IP.Equal(state.IP) {
		updateDTO.IP = plan.IP.ValueString()
	}
	if !plan.Description.Equal(state.Description) {
		updateDTO.Description = plan.Description.ValueString()
	}
	if !plan.Port.Equal(state.Port) {
		updateDTO.Port = int(plan.Port.ValueInt32())
	}
	if !plan.User.Equal(state.User) {
		updateDTO.User = plan.User.ValueString()
	}
	if !plan.PrivateKeyUUID.Equal(state.PrivateKeyUUID) {
		updateDTO.PrivateKeyUUID = plan.PrivateKeyUUID.ValueString()
	}

	err := r.client.Server.Update(state.ID.ValueString(), &updateDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update server, got error: %s", err))
		return
	}

	server, err := r.client.Server.Get(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read server after update, got error: %s", err))
		return
	}

	var newState ServerModel

	mapCommonServerFields(&newState, server)
	newState.PrivateKeyUUID = plan.PrivateKeyUUID

	diags = resp.State.Set(ctx, &newState)
	resp.Diagnostics.Append(diags...)
}
