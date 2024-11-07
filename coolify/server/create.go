package server

import (
	"context"
	"fmt"

	coolify_sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (r *ServerResource) CreateServer(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServerModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Planned data before creation", map[string]interface{}{
		"name":        data.Name.ValueString(),
		"ip":          data.IP.ValueString(),
		"description": data.Description.ValueString(),
	})

	privateKeyUUID := data.PrivateKeyUUID.ValueString()

	createDTO := coolify_sdk.CreateServerDTO{
		Name:           data.Name.ValueString(),
		IP:             data.IP.ValueString(),
		Description:    data.Description.ValueString(),
		IsBuildServer:  false,
		Port:           int(data.Port.ValueInt32()),
		User:           data.User.ValueString(),
		PrivateKeyUUID: data.PrivateKeyUUID.ValueString(),
	}

	serverID, err := r.client.Server.Create(&createDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create server, got error: %s", err))
		return
	}

	data.ID = types.StringValue(*serverID)
	data.PrivateKeyUUID = types.StringValue(privateKeyUUID)
	tflog.Trace(ctx, "Created a server", map[string]interface{}{"server_id": serverID})

	tflog.Debug(ctx, "Data after server creation", map[string]interface{}{
		"uuid":        data.ID.ValueString(),
		"name":        data.Name.ValueString(),
		"ip":          data.IP.ValueString(),
		"description": data.Description.ValueString(),
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	tflog.Debug(ctx, "Server state saved to file after creation")
}
