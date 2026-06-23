package application_restart

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/marconneves/coolify-sdk-go/application"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// ApplicationRestartModel describes the Terraform resource state for a restart
// operation on a Coolify application.
type ApplicationRestartModel struct {
	ApplicationUUID types.String `tfsdk:"application_uuid"`
	Message         types.String `tfsdk:"message"`
	DeploymentUUID  types.String `tfsdk:"deployment_uuid"`
}

func mapRestartResponseToModel(data *ApplicationRestartModel, resp *application.RestartResponse) {
	data.Message = types.StringValue(resp.Message)
	data.DeploymentUUID = types.StringValue(resp.DeploymentUUID)
}

func mapApplicationStatusToModel(data *ApplicationRestartModel, app *application.Application) {
	data.Message = configure.ValueStringValue(&app.Status, data.Message)
}
