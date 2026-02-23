package database_mysql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// MySQLDataSource represents a Coolify MySQL database data source.
var _ datasource.DataSource = &MySQLDataSource{}

// NewMySQLDataSource creates a new MySQL database data source.
func NewMySQLDataSource() datasource.DataSource {
	return &MySQLDataSource{}
}

// MySQLDataSource represents the data source implementation.
type MySQLDataSource struct {
	client *coolify_sdk.Sdk
}

// Metadata returns the data source metadata.
func (d *MySQLDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_database_mysql"
}

// Schema returns the data source schema.
func (d *MySQLDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Get a Coolify MySQL database",

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
			
			"mysql_user": schema.StringAttribute{
				MarkdownDescription: "MySQL user",
				Computed:            true,
			},
			"mysql_database": schema.StringAttribute{
				MarkdownDescription: "MySQL database name",
				Computed:            true,
			},
			"mysql_conf": schema.StringAttribute{
				MarkdownDescription: "MySQL configuration",
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
func (d *MySQLDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	configure.ConfigureClient(ctx, req, &d.client)
}

// Read retrieves information for a MySQL database.
func (d *MySQLDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DatabaseMySQLModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	db, err := d.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read MySQL database, got error: %s", err))
		return
	}

	mapMySQLResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
