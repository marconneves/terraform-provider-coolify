package application_start

import (
	"context"
	"fmt"

	"github.com/marconneves/coolify-sdk-go/application"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

// StartApplication calls the Coolify API to start an application and
// updates the model in-place with the response.
func StartApplication(ctx context.Context, client *coolify_sdk.Sdk, data *ApplicationStartModel) diag.Diagnostics {
	var diags diag.Diagnostics

	uuid := data.ApplicationUUID.ValueString()

	opts := &application.StartOptions{}
	if !data.Force.IsNull() {
		force := data.Force.ValueBool()
		opts.Force = &force
	}
	if !data.InstantDeploy.IsNull() {
		instantDeploy := data.InstantDeploy.ValueBool()
		opts.InstantDeploy = &instantDeploy
	}

	tflog.Debug(ctx, "Starting application", map[string]interface{}{
		"application_uuid": uuid,
		"force":            opts.Force,
		"instant_deploy":   opts.InstantDeploy,
	})

	startResp, err := client.Application.Start(ctx, uuid, opts)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to start application, got error: %s", err))
		return diags
	}

	mapStartResponseToModel(data, startResp)

	tflog.Debug(ctx, "Application started successfully", map[string]interface{}{
		"application_uuid": uuid,
		"deployment_uuid":  startResp.DeploymentUUID,
	})

	return diags
}
