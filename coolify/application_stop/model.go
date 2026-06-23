package application_stop

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/marconneves/coolify-sdk-go/application"
)

// ApplicationStopModel describes the Terraform resource state for a stop
// operation on a Coolify application.
type ApplicationStopModel struct {
	ApplicationUUID types.String `tfsdk:"application_uuid"`
	DockerCleanup   types.Bool   `tfsdk:"docker_cleanup"`
}

func mapApplicationStatusToModel(data *ApplicationStopModel, app *application.Application) {
	// Status is purely informational after the stop action.
}
