package server

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ resource.Resource = &ServerResource{}
var _ resource.ResourceWithImportState = &ServerResource{}

func NewServerResource() resource.Resource {
	return &ServerResource{}
}

type ServerResource struct {
	client *coolify_sdk.Sdk
}

func (r *ServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (r *ServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage Coolify servers",

		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				MarkdownDescription: "Server identifier",
				Computed:            true,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Server name",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Server description",
				Optional:            true,
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: "Server IP address",
				Required:            true,
			},
			"port": schema.Int32Attribute{
				MarkdownDescription: "Server port",
				Required:            true,
			},
			"user": schema.StringAttribute{
				MarkdownDescription: "Server user",
				Required:            true,
			},
			"private_key_uuid": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The UUID of the private key associated with the server.",
			},
		},
	}
}

func (r *ServerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &r.client)
}

func (r *ServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.CreateServer(ctx, req, resp)
}

func (r *ServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.ReadServerResource(ctx, req, resp)
}

func (r *ServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.UpdateServer(ctx, req, resp)
}

func (r *ServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.DeleteServer(ctx, req, resp)
}

func (r *ServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
