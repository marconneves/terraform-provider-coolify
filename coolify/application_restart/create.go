package application_restart

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

// RestartApplication calls the Coolify API to restart an application and
// updates the model in-place with the response.
func RestartApplication(ctx context.Context, client *coolify_sdk.Sdk, data *ApplicationRestartModel) diag.Diagnostics {
	var diags diag.Diagnostics

	uuid := data.ApplicationUUID.ValueString()

	tflog.Debug(ctx, "Restarting application", map[string]interface{}{
		"application_uuid": uuid,
	})

	restartResp, err := client.Application.Restart(ctx, uuid)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to restart application, got error: %s", err))
		return diags
	}

	mapRestartResponseToModel(data, restartResp)

	tflog.Debug(ctx, "Application restarted successfully", map[string]interface{}{
		"application_uuid": uuid,
		"deployment_uuid":  restartResp.DeploymentUUID,
	})

	return diags
}
