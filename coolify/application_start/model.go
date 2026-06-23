package application_start

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/marconneves/coolify-sdk-go/application"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// ApplicationStartModel describes the Terraform resource state for a start
// operation on a Coolify application.
type ApplicationStartModel struct {
	ApplicationUUID types.String `tfsdk:"application_uuid"`
	Force           types.Bool   `tfsdk:"force"`
	InstantDeploy   types.Bool   `tfsdk:"instant_deploy"`
	Message         types.String `tfsdk:"message"`
	DeploymentUUID  types.String `tfsdk:"deployment_uuid"`
}

func mapStartResponseToModel(data *ApplicationStartModel, resp *application.StartResponse) {
	data.Message = types.StringValue(resp.Message)
	data.DeploymentUUID = types.StringValue(resp.DeploymentUUID)
}

func mapApplicationStatusToModel(data *ApplicationStartModel, app *application.Application) {
	data.Message = configure.ValueStringValue(&app.Status, data.Message)
}
