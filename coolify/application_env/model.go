package application_env

import "github.com/hashicorp/terraform-plugin-framework/types"

// ApplicationEnvModel describes a single environment variable managed as its
// own Terraform resource, allowing the use of flags such as `is_preview`,
// `is_literal`, `is_multiline` and `is_shown_once`.
type ApplicationEnvModel struct {
	Id              types.String `tfsdk:"id"`
	ApplicationUUID types.String `tfsdk:"application_uuid"`
	Key             types.String `tfsdk:"key"`
	Value           types.String `tfsdk:"value"`
	IsPreview       types.Bool   `tfsdk:"is_preview"`
	IsLiteral       types.Bool   `tfsdk:"is_literal"`
	IsMultiline     types.Bool   `tfsdk:"is_multiline"`
	IsShownOnce     types.Bool   `tfsdk:"is_shown_once"`
}
