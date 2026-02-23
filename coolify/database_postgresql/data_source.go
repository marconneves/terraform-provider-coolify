package database_postgresql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// PostgresDataSource represents a Coolify PostgreSQL database data source.
var _ datasource.DataSource = &PostgresDataSource{}

// NewPostgresDataSource creates a new PostgreSQL database data source.
func NewPostgresDataSource() datasource.DataSource {
	return &PostgresDataSource{}
}

// PostgresDataSource represents the data source implementation.
type PostgresDataSource struct {
	client *coolify_sdk.Sdk
}

// Metadata returns the data source metadata.
func (d *PostgresDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_database_postgresql"
}

// Schema returns the data source schema.
func (d *PostgresDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Get a Coolify PostgreSQL database",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Database identifier",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Database name",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Database description",
				Computed:            true,
			},
			"image": schema.StringAttribute{
				MarkdownDescription: "Database image",
				Computed:            true,
			},
			"is_public": schema.BoolAttribute{
				MarkdownDescription: "Is the database public",
				Computed:            true,
			},
			"public_port": schema.Int64Attribute{
				MarkdownDescription: "Public port",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Database status",
				Computed:            true,
			},
			
			"postgres_user": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL user",
				Computed:            true,
			},
			"postgres_db": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL database name",
				Computed:            true,
			},
			"postgres_initdb_args": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL initdb args",
				Computed:            true,
			},
			"postgres_host_auth_method": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL host auth method",
				Computed:            true,
			},
			"postgres_conf": schema.StringAttribute{
				MarkdownDescription: "PostgreSQL configuration",
				Computed:            true,
			},

			"limits_memory": schema.StringAttribute{
				MarkdownDescription: "Memory limit",
				Computed:            true,
			},
			"limits_memory_swap": schema.StringAttribute{
				MarkdownDescription: "Memory swap limit",
				Computed:            true,
			},
			"limits_memory_swappiness": schema.Int64Attribute{
				MarkdownDescription: "Memory swappiness",
				Computed:            true,
			},
			"limits_memory_reservation": schema.StringAttribute{
				MarkdownDescription: "Memory reservation limit",
				Computed:            true,
			},
			"limits_cpus": schema.StringAttribute{
				MarkdownDescription: "CPUs limit",
				Computed:            true,
			},
			"limits_cpuset": schema.StringAttribute{
				MarkdownDescription: "CPUs set limit",
				Computed:            true,
			},
			"limits_cpu_shares": schema.Int64Attribute{
				MarkdownDescription: "CPU shares limit",
				Computed:            true,
			},
		},
	}
}

// Configure configures the data source.
func (d *PostgresDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &d.client)
}

// Read retrieves information for a PostgreSQL database.
func (d *PostgresDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DatabasePostgresModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	db, err := d.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read PostgreSQL database, got error: %s", err))
		return
	}

	mapPostgresResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
