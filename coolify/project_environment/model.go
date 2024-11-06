package project_environment

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

type EnvironmentModel struct {
	Id          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
	ProjectUUID types.String `tfsdk:"project_uuid"`
	ProjectID   types.Int64  `tfsdk:"project_id"`
}

func mapEnvironmentModel(data *EnvironmentModel, environment *coolify_sdk.EnvironmentData) {
	data.Id = types.Int64Value(int64(environment.Id))
	data.Name = types.StringValue(environment.Name)
	data.Description = types.StringValue(environment.Description)
	data.ProjectID = types.Int64Value(int64(environment.ProjectID))
	data.CreatedAt = types.StringValue(environment.CreatedAt.String())
	data.UpdatedAt = types.StringValue(environment.UpdatedAt.String())

}
