package private_key

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
)

type PrivateKeyModel struct {
	ID           types.Int64  `tfsdk:"id"`
	UUID         types.String `tfsdk:"uuid"`
	Name         types.String `tfsdk:"name"`
	Description  types.String `tfsdk:"description"`
	PrivateKey   types.String `tfsdk:"private_key"`
	IsGitRelated types.Bool   `tfsdk:"is_git_related"`
	TeamID       types.Int64  `tfsdk:"team_id"`
	CreatedAt    types.String `tfsdk:"created_at"`
	UpdatedAt    types.String `tfsdk:"updated_at"`
}

func mapPrivateKeyModel(data *PrivateKeyModel, privateKey *coolify_sdk.PrivateKey) {
	data.ID = types.Int64Value(int64(privateKey.ID))
	data.UUID = types.StringValue(privateKey.UUID)
	data.Name = types.StringValue(privateKey.Name)
	data.Description = types.StringValue(privateKey.Description)
	data.PrivateKey = types.StringValue(privateKey.PrivateKey)
	data.IsGitRelated = types.BoolValue(privateKey.IsGitRelated)
	data.TeamID = types.Int64Value(int64(privateKey.TeamID))
	data.CreatedAt = types.StringValue(privateKey.CreatedAt.String())
	data.UpdatedAt = types.StringValue(privateKey.UpdatedAt.String())
}
