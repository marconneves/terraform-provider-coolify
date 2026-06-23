package application_restart

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ resource.Resource = &ApplicationRestartResource{}
var _ resource.ResourceWithImportState = &ApplicationRestartResource{}

func NewApplicationRestartResource() resource.Resource {
	return &ApplicationRestartResource{}
}

type ApplicationRestartResource struct {
	client *coolify_sdk.Sdk
}

func (r *ApplicationRestartResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_restart"
}

func (r *ApplicationRestartResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Restart a Coolify application. Creating this resource sends a restart signal to the specified application. Deleting this resource is a no-op — it does not affect the application.",

		Attributes: map[string]schema.Attribute{
			"application_uuid": schema.StringAttribute{
				MarkdownDescription: "UUID of the application to restart.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"message": schema.StringAttribute{
				MarkdownDescription: "Status message returned by the Coolify API after the restart action.",
				Computed:            true,
			},
			"deployment_uuid": schema.StringAttribute{
				MarkdownDescription: "UUID of the deployment triggered by the restart action.",
				Computed:            true,
			},
		},
	}
}

func (r *ApplicationRestartResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &r.client)
}

func (r *ApplicationRestartResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApplicationRestartModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags := r.restart(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationRestartResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.ReadApplicationRestart(ctx, req, resp)
}

func (r *ApplicationRestartResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ApplicationRestartModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags := r.restart(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationRestartResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Deleting the resource does not affect the application — it's a no-op.
}

func (r *ApplicationRestartResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("application_uuid"), req, resp)
}

func (r *ApplicationRestartResource) restart(ctx context.Context, data *ApplicationRestartModel) diag.Diagnostics {
	return RestartApplication(ctx, r.client, data)
}
