package application_start

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

var _ resource.Resource = &ApplicationStartResource{}
var _ resource.ResourceWithImportState = &ApplicationStartResource{}

func NewApplicationStartResource() resource.Resource {
	return &ApplicationStartResource{}
}

type ApplicationStartResource struct {
	client *coolify_sdk.Sdk
}

func (r *ApplicationStartResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_start"
}

func (r *ApplicationStartResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Start a Coolify application. Creating this resource sends a start signal to the specified application. Deleting this resource is a no-op — it does not stop the application.",

		Attributes: map[string]schema.Attribute{
			"application_uuid": schema.StringAttribute{
				MarkdownDescription: "UUID of the application to start.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"force": schema.BoolAttribute{
				MarkdownDescription: "Force restart even if the application is already running.",
				Optional:            true,
			},
			"instant_deploy": schema.BoolAttribute{
				MarkdownDescription: "Deploy immediately when starting.",
				Optional:            true,
			},
			"message": schema.StringAttribute{
				MarkdownDescription: "Status message returned by the Coolify API after the start action.",
				Computed:            true,
			},
			"deployment_uuid": schema.StringAttribute{
				MarkdownDescription: "UUID of the deployment triggered by the start action.",
				Computed:            true,
			},
		},
	}
}

func (r *ApplicationStartResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &r.client)
}

func (r *ApplicationStartResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApplicationStartModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags := r.start(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationStartResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.ReadApplicationStart(ctx, req, resp)
}

func (r *ApplicationStartResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ApplicationStartModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags := r.start(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ApplicationStartResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Deleting the resource does not stop the application — it's a no-op.
}

func (r *ApplicationStartResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("application_uuid"), req, resp)
}

// start performs the application start action and updates the model with the response.
func (r *ApplicationStartResource) start(ctx context.Context, data *ApplicationStartModel) diag.Diagnostics {
	return StartApplication(ctx, r.client, data)
}
