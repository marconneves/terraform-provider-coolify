package application_dockerimage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ resource.Resource = &ApplicationDockerImageResource{}

func NewApplicationDockerImageResource() resource.Resource {
	return &ApplicationDockerImageResource{}
}

type ApplicationDockerImageResource struct {
	client *coolify_sdk.Sdk
}

func (r *ApplicationDockerImageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_dockerimage"
}

func (r *ApplicationDockerImageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Deploy a Coolify application based on a prebuilt Docker image hosted on any registry (Docker Hub, GHCR, GCP Artifact Registry, AWS ECR, etc.). The `environment_variables` map performs a bulk upsert against the application after creation, mirroring the ergonomics of `google_cloud_run_v2_service.template[0].containers[0].env`.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Application UUID assigned by Coolify.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Application name.",
				Optional:            true,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Application description.",
				Optional:            true,
				Computed:            true,
			},
			"project_uuid": schema.StringAttribute{
				MarkdownDescription: "Coolify project UUID.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"server_uuid": schema.StringAttribute{
				MarkdownDescription: "Coolify server UUID.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"environment_name": schema.StringAttribute{
				MarkdownDescription: "Environment name. Provide either `environment_name` or `environment_uuid`.",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"environment_uuid": schema.StringAttribute{
				MarkdownDescription: "Environment UUID. Provide either `environment_name` or `environment_uuid`.",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"destination_uuid": schema.StringAttribute{
				MarkdownDescription: "Destination UUID. Required when the server has multiple destinations.",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},

			"docker_registry_image_name": schema.StringAttribute{
				MarkdownDescription: "Fully qualified image name. For GCP Artifact Registry use the `LOCATION-docker.pkg.dev/PROJECT/REPO/IMAGE` form.",
				Required:            true,
			},
			"docker_registry_image_tag": schema.StringAttribute{
				MarkdownDescription: "Image tag. Defaults to `latest` on the Coolify side when omitted.",
				Optional:            true,
				Computed:            true,
			},

			"ports_exposes": schema.StringAttribute{
				MarkdownDescription: "Container ports to expose, comma-separated (e.g. `\"8080\"` or `\"80,443\"`).",
				Required:            true,
			},
			"ports_mappings": schema.StringAttribute{
				MarkdownDescription: "Host-to-container port mappings, comma-separated (e.g. `\"8080:80\"`).",
				Optional:            true,
				Computed:            true,
			},
			"domains": schema.StringAttribute{
				MarkdownDescription: "Application URLs in a comma-separated list.",
				Optional:            true,
				Computed:            true,
			},

			"instant_deploy": schema.BoolAttribute{
				MarkdownDescription: "Deploy immediately after creation.",
				Optional:            true,
			},
			"is_force_https_enabled": schema.BoolAttribute{
				MarkdownDescription: "Force HTTPS.",
				Optional:            true,
				Computed:            true,
			},
			"connect_to_docker_network": schema.BoolAttribute{
				MarkdownDescription: "Connect the application to Coolify's predefined Docker network.",
				Optional:            true,
				Computed:            true,
			},

			"health_check_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable health checks.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_path": schema.StringAttribute{
				MarkdownDescription: "Health check path.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_port": schema.StringAttribute{
				MarkdownDescription: "Health check port.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_host": schema.StringAttribute{
				MarkdownDescription: "Health check host.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_method": schema.StringAttribute{
				MarkdownDescription: "Health check HTTP method.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_return_code": schema.Int64Attribute{
				MarkdownDescription: "Expected health check HTTP status code.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_scheme": schema.StringAttribute{
				MarkdownDescription: "Health check scheme (`http` or `https`).",
				Optional:            true,
				Computed:            true,
			},
			"health_check_response_text": schema.StringAttribute{
				MarkdownDescription: "Health check expected response body.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_interval": schema.Int64Attribute{
				MarkdownDescription: "Health check interval in seconds.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_timeout": schema.Int64Attribute{
				MarkdownDescription: "Health check timeout in seconds.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_retries": schema.Int64Attribute{
				MarkdownDescription: "Health check retries.",
				Optional:            true,
				Computed:            true,
			},
			"health_check_start_period": schema.Int64Attribute{
				MarkdownDescription: "Health check start period in seconds.",
				Optional:            true,
				Computed:            true,
			},

			"limits_memory": schema.StringAttribute{
				MarkdownDescription: "Memory limit.",
				Optional:            true,
				Computed:            true,
			},
			"limits_memory_swap": schema.StringAttribute{
				MarkdownDescription: "Memory swap limit.",
				Optional:            true,
				Computed:            true,
			},
			"limits_memory_swappiness": schema.Int64Attribute{
				MarkdownDescription: "Memory swappiness.",
				Optional:            true,
				Computed:            true,
			},
			"limits_memory_reservation": schema.StringAttribute{
				MarkdownDescription: "Memory reservation.",
				Optional:            true,
				Computed:            true,
			},
			"limits_cpus": schema.StringAttribute{
				MarkdownDescription: "CPU limit.",
				Optional:            true,
				Computed:            true,
			},
			"limits_cpuset": schema.StringAttribute{
				MarkdownDescription: "CPU set.",
				Optional:            true,
				Computed:            true,
			},
			"limits_cpu_shares": schema.Int64Attribute{
				MarkdownDescription: "CPU shares.",
				Optional:            true,
				Computed:            true,
			},

			"custom_labels": schema.StringAttribute{
				MarkdownDescription: "Custom Docker container labels.",
				Optional:            true,
				Computed:            true,
			},
			"custom_docker_run_options": schema.StringAttribute{
				MarkdownDescription: "Custom `docker run` options.",
				Optional:            true,
				Computed:            true,
			},

			"environment_variables": schema.MapAttribute{
				MarkdownDescription: "Environment variables to upsert in bulk on the application. Variables not listed here are left untouched, so individual entries can be managed via `coolify_application_env`.",
				Optional:            true,
				ElementType:         types.StringType,
			},

			"status": schema.StringAttribute{
				MarkdownDescription: "Application status reported by Coolify.",
				Computed:            true,
			},
			"fqdn": schema.StringAttribute{
				MarkdownDescription: "Application fully qualified domain name(s) as a comma-separated string.",
				Computed:            true,
			},
		},
	}
}

func (r *ApplicationDockerImageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &r.client)
}

func (r *ApplicationDockerImageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.CreateApplication(ctx, req, resp)
}

func (r *ApplicationDockerImageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.ReadApplication(ctx, req, resp)
}

func (r *ApplicationDockerImageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.UpdateApplication(ctx, req, resp)
}

func (r *ApplicationDockerImageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.DeleteApplication(ctx, req, resp)
}
