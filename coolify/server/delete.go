package server

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (r *ServerResource) DeleteServer(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServerModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.Server.Delete(data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete server, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted server", map[string]interface{}{"server_id": data.ID.ValueString()})
}
