package application_env

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	"github.com/marconneves/coolify-sdk-go/application"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ resource.Resource = &ApplicationEnvResource{}

func NewApplicationEnvResource() resource.Resource {
	return &ApplicationEnvResource{}
}

type ApplicationEnvResource struct {
	client *coolify_sdk.Sdk
}

func (r *ApplicationEnvResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_env"
}

func (r *ApplicationEnvResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage a single environment variable on a Coolify application. Use this resource when you need to set per-variable flags such as `is_preview`, `is_literal`, `is_multiline` or `is_shown_once`. For the simple `KEY = value` case, prefer the `environment_variables` map on `coolify_application_dockerimage`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Environment variable UUID.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"application_uuid": schema.StringAttribute{
				MarkdownDescription: "UUID of the application this variable belongs to.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"key": schema.StringAttribute{
				MarkdownDescription: "Environment variable name.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "Environment variable value.",
				Required:            true,
				Sensitive:           true,
			},
			"is_preview": schema.BoolAttribute{
				MarkdownDescription: "Whether the variable is used in preview deployments.",
				Optional:            true,
				Computed:            true,
			},
			"is_literal": schema.BoolAttribute{
				MarkdownDescription: "Whether the value is treated as a literal (no escaping).",
				Optional:            true,
				Computed:            true,
			},
			"is_multiline": schema.BoolAttribute{
				MarkdownDescription: "Whether the value spans multiple lines.",
				Optional:            true,
				Computed:            true,
			},
			"is_shown_once": schema.BoolAttribute{
				MarkdownDescription: "Whether the value is shown only once in the Coolify UI.",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (r *ApplicationEnvResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &r.client)
}

func (r *ApplicationEnvResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApplicationEnvModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := application.EnvDTO{
		Key:         data.Key.ValueString(),
		Value:       data.Value.ValueString(),
		IsPreview:   data.IsPreview.ValueBoolPointer(),
		IsLiteral:   data.IsLiteral.ValueBoolPointer(),
		IsMultiline: data.IsMultiline.ValueBoolPointer(),
		IsShownOnce: data.IsShownOnce.ValueBoolPointer(),
	}

	if _, err := r.client.Application.CreateEnv(ctx, data.ApplicationUUID.ValueString(), &dto); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create application env, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(refreshFromAPI(ctx, r, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationEnvResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApplicationEnvModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshFromAPI(ctx, r, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationEnvResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ApplicationEnvModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := application.EnvDTO{
		Key:         data.Key.ValueString(),
		Value:       data.Value.ValueString(),
		IsPreview:   data.IsPreview.ValueBoolPointer(),
		IsLiteral:   data.IsLiteral.ValueBoolPointer(),
		IsMultiline: data.IsMultiline.ValueBoolPointer(),
		IsShownOnce: data.IsShownOnce.ValueBoolPointer(),
	}

	if _, err := r.client.Application.UpdateEnv(ctx, data.ApplicationUUID.ValueString(), &dto); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update application env, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(refreshFromAPI(ctx, r, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationEnvResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApplicationEnvModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() || data.Id.ValueString() == "" {
		return
	}

	if err := r.client.Application.DeleteEnv(ctx, data.ApplicationUUID.ValueString(), data.Id.ValueString()); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete application env, got error: %s", err))
		return
	}
}

// refreshFromAPI fetches the env list for the application and updates the
// state for the env matching `data.Key`.
func refreshFromAPI(ctx context.Context, r *ApplicationEnvResource, data *ApplicationEnvModel) diag.Diagnostics {
	var diags diag.Diagnostics

	envs, err := r.client.Application.ListEnvs(ctx, data.ApplicationUUID.ValueString())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to list envs for application %s: %s", data.ApplicationUUID.ValueString(), err))
		return diags
	}

	for _, e := range *envs {
		if e.Key != data.Key.ValueString() {
			continue
		}
		data.Id = types.StringValue(e.UUID)
		data.Value = types.StringValue(e.Value)
		data.IsPreview = types.BoolValue(e.IsPreview)
		data.IsLiteral = types.BoolValue(e.IsLiteral)
		data.IsMultiline = types.BoolValue(e.IsMultiline)
		data.IsShownOnce = types.BoolValue(e.IsShownOnce)
		return diags
	}

	diags.AddError("Client Error", fmt.Sprintf("Env %q not found on application %s after API call", data.Key.ValueString(), data.ApplicationUUID.ValueString()))
	return diags
}
