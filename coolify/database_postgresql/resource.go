package database_postgresql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ resource.Resource = &PostgresResource{}

func NewPostgresResource() resource.Resource {
	return &PostgresResource{}
}

type PostgresResource struct {
	client *coolify_sdk.Sdk
}

func (r *PostgresResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_database_postgresql"
}

func (r *PostgresResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage Coolify PostgreSQL databases",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Database identifier",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Database name",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Database description",
				Optional:            true,
			},
			"server_uuid": schema.StringAttribute{
				MarkdownDescription: "Server UUID",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"project_uuid": schema.StringAttribute{
				MarkdownDescription: "Project UUID",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"environment_name": schema.StringAttribute{
				MarkdownDescription: "Environment name",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"destination_uuid": schema.StringAttribute{
				MarkdownDescription: "Destination UUID",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"image": schema.StringAttribute{
				MarkdownDescription: "Database image",
				Optional:            true,
				Computed:            true,
			},
			"is_public": schema.BoolAttribute{
				MarkdownDescription: "Is the database public",
				Optional:            true,
				Computed:            true,
			},
			"public_port": schema.Int64Attribute{
				MarkdownDescription: "Public port",
				Optional:            true,
				Computed:            true,
			},
			"instant_deploy": schema.BoolAttribute{
				MarkdownDescription: "Instant deploy",
				Optional:            true,
			},

			"postgres_user": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL user",
				Optional:            true,
				Computed:            true,
			},
			"postgres_password": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL password",
				Optional:            true,
				Computed:            true,
				Sensitive:           true,
			},
			"postgres_db": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL database name",
				Optional:            true,
				Computed:            true,
			},
			"postgres_initdb_args": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL initdb args",
				Optional:            true,
				Computed:            true,
			},
			"postgres_host_auth_method": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL host auth method",
				Optional:            true,
				Computed:            true,
			},
			"postgres_conf": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL configuration",
				Optional:            true,
				Computed:            true,
			},

			"limits_memory": schema.StringAttribute{
				MarkdownDescription: "Memory limit",
				Optional:            true,
				Computed:            true,
			},
			"limits_memory_swap": schema.StringAttribute{
				MarkdownDescription: "Memory swap limit",
				Optional:            true,
				Computed:            true,
			},
			"limits_memory_swappiness": schema.Int64Attribute{
				MarkdownDescription: "Memory swappiness",
				Optional:            true,
				Computed:            true,
			},
			"limits_memory_reservation": schema.StringAttribute{
				MarkdownDescription: "Memory reservation limit",
				Optional:            true,
				Computed:            true,
			},
			"limits_cpus": schema.StringAttribute{
				MarkdownDescription: "CPUs limit",
				Optional:            true,
				Computed:            true,
			},
			"limits_cpuset": schema.StringAttribute{
				MarkdownDescription: "CPUs set limit",
				Optional:            true,
				Computed:            true,
			},
			"limits_cpu_shares": schema.Int64Attribute{
				MarkdownDescription: "CPU shares limit",
				Optional:            true,
				Computed:            true,
			},

			"status": schema.StringAttribute{
				MarkdownDescription: "Database status",
				Computed:            true,
			},
		},
	}
}

func (r *PostgresResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &r.client)
}

func (r *PostgresResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.CreatePostgres(ctx, req, resp)
}

func (r *PostgresResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.ReadPostgres(ctx, req, resp)
}

func (r *PostgresResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.UpdatePostgres(ctx, req, resp)
}

func (r *PostgresResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.DeletePostgres(ctx, req, resp)
}
