package application_start

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (r *ApplicationStartResource) ReadApplicationStart(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApplicationStartModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	uuid := data.ApplicationUUID.ValueString()

	app, err := r.client.Application.Get(ctx, uuid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read application, got error: %s", err))
		return
	}

	mapApplicationStatusToModel(&data, app)

	tflog.Debug(ctx, "Read application start state", map[string]interface{}{
		"application_uuid": uuid,
		"status":           app.Status,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
