package project_environment

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

func readEnvironment(client coolify_sdk.Sdk, projectUUID types.String, environmentName types.String) (*coolify_sdk.EnvironmentData, diag.Diagnostics) {
	var diags diag.Diagnostics

	if projectUUID.IsNull() || environmentName.IsNull() {
		diags.AddError("Configuration Error", "Both 'project_uuid' and 'environment_name' must be specified.")
		return nil, diags
	}

	environment, err := client.Project.Environment(projectUUID.ValueString(), environmentName.ValueString())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read environment data, got error: %s", err))
		return nil, diags
	}

	if environment == nil {
		diags.AddError("Not Found", "No environment found with the specified name and project UUID")
		return nil, diags
	}

	return environment, diags
}
