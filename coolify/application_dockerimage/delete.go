package application_dockerimage

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func (r *ApplicationDockerImageResource) DeleteApplication(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApplicationDockerImageModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.Application.Delete(ctx, data.Id.ValueString(), nil); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete docker image application, got error: %s", err))
		return
	}
}
