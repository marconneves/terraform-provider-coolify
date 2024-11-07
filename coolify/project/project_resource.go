package project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ resource.Resource = &ProjectResource{}
var _ resource.ResourceWithImportState = &ProjectResource{}

func NewProjectResource() resource.Resource {
	return &ProjectResource{}
}

type ProjectResource struct {
	client *coolify_sdk.Sdk
}

func (r *ProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (r *ProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage Coolify projects",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Project identifier",
				Computed:            true,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Project name",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Project description",
				Optional:            true,
			},
		},
	}
}

func (r *ProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &r.client)
}

func (r *ProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.CreateProject(ctx, req, resp)
}

func (r *ProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.ReadProjectResource(ctx, req, resp)
}

func (r *ProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.UpdateProject(ctx, req, resp)
}

func (r *ProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.DeleteProject(ctx, req, resp)
}

func (r *ProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
