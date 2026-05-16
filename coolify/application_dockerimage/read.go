package application_dockerimage

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func (r *ApplicationDockerImageResource) ReadApplication(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApplicationDockerImageModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	app, err := r.client.Application.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read docker image application, got error: %s", err))
		return
	}

	mapApplicationToModel(&data, app)

	envs, err := r.client.Application.ListEnvs(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list envs, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(mapEnvsToModel(ctx, &data, *envs)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
