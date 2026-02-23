package database_mariadb

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// MariaDBDataSource represents a Coolify MariaDB database data source.
var _ datasource.DataSource = &MariaDBDataSource{}

// NewMariaDBDataSource creates a new MariaDB database data source.
func NewMariaDBDataSource() datasource.DataSource {
	return &MariaDBDataSource{}
}

// MariaDBDataSource represents the data source implementation.
type MariaDBDataSource struct {
	client *coolify_sdk.Sdk
}

// Metadata returns the data source metadata.
func (d *MariaDBDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_database_mariadb"
}

// Schema returns the data source schema.
func (d *MariaDBDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Get a Coolify MariaDB database",

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
			
			"mariadb_user": schema.StringAttribute{
				MarkdownDescription: "MariaDB user",
				Computed:            true,
			},
			"mariadb_database": schema.StringAttribute{
				MarkdownDescription: "MariaDB database name",
				Computed:            true,
			},
			"mariadb_conf": schema.StringAttribute{
				MarkdownDescription: "MariaDB configuration",
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
func (d *MariaDBDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &d.client)
}

// Read retrieves information for a MariaDB database.
func (d *MariaDBDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DatabaseMariaDBModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	db, err := d.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read MariaDB database, got error: %s", err))
		return
	}

	mapMariaDBResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
