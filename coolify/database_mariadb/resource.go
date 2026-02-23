package database_mariadb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// MariaDBResource defines the resource implementation.
var _ resource.Resource = &MariaDBResource{}

// NewMariaDBResource creates a new MariaDB database resource.
func NewMariaDBResource() resource.Resource {
	return &MariaDBResource{}
}

// MariaDBResource represents the resource implementation.
type MariaDBResource struct {
	client *coolify_sdk.Sdk
}

// Metadata returns the resource metadata.
func (r *MariaDBResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_database_mariadb"
}

// Schema returns the resource schema.
func (r *MariaDBResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manage Coolify MariaDB databases",

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
			"environment_uuid": schema.StringAttribute{
				MarkdownDescription: "Environment UUID",
				Optional:            true,
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

			"mariadb_root_password": schema.StringAttribute{
				MarkdownDescription: "MariaDB root password",
				Optional:            true,
				Computed:            true,
				Sensitive:           true,
			},
			"mariadb_password": schema.StringAttribute{
				MarkdownDescription: "MariaDB password",
				Optional:            true,
				Computed:            true,
				Sensitive:           true,
			},
			"mariadb_user": schema.StringAttribute{
				MarkdownDescription: "MariaDB user",
				Optional:            true,
				Computed:            true,
			},
			"mariadb_database": schema.StringAttribute{
				MarkdownDescription: "MariaDB database name",
				Optional:            true,
				Computed:            true,
			},
			"mariadb_conf": schema.StringAttribute{
				MarkdownDescription: "MariaDB configuration",
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

// Configure configures the resource.
func (r *MariaDBResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &r.client)
}

// Create creates a new MariaDB database instance.
func (r *MariaDBResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	r.CreateMariaDB(ctx, req, resp)
}

// Read reads a MariaDB database instance.
func (r *MariaDBResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	r.ReadMariaDB(ctx, req, resp)
}

// Update updates a MariaDB database instance.
func (r *MariaDBResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	r.UpdateMariaDB(ctx, req, resp)
}

// Delete deletes a MariaDB database instance.
func (r *MariaDBResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.DeleteMariaDB(ctx, req, resp)
}
