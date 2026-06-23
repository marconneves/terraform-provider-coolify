package application_stop

import (
	"context"
	"fmt"

	"github.com/marconneves/coolify-sdk-go/application"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

// StopApplication calls the Coolify API to stop an application.
func StopApplication(ctx context.Context, client *coolify_sdk.Sdk, data *ApplicationStopModel) diag.Diagnostics {
	var diags diag.Diagnostics

	uuid := data.ApplicationUUID.ValueString()

	opts := &application.StopOptions{}
	if !data.DockerCleanup.IsNull() {
		dockerCleanup := data.DockerCleanup.ValueBool()
		opts.DockerCleanup = &dockerCleanup
	}

	tflog.Debug(ctx, "Stopping application", map[string]interface{}{
		"application_uuid": uuid,
		"docker_cleanup":   opts.DockerCleanup,
	})

	err := client.Application.Stop(ctx, uuid, opts)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to stop application, got error: %s", err))
		return diags
	}

	tflog.Debug(ctx, "Application stopped successfully", map[string]interface{}{
		"application_uuid": uuid,
	})

	return diags
}
