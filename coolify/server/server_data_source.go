package server

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

var _ datasource.DataSource = &ServerDataSource{}

func NewServerDataSource() datasource.DataSource {
	return &ServerDataSource{}
}

type ServerDataSource struct {
	client *coolify_sdk.Sdk
}

func (s *ServerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_server"
}

func (s *ServerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"uuid": schema.StringAttribute{
				Optional:    true,
				Description: "The unique identifier of the server.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Description: "The name of the server.",
			},
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "The IP address of the server.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "A description of the server.",
			},
			"high_disk_usage_notification_sent": schema.BoolAttribute{
				Computed:    true,
				Description: "Indicates if a high disk usage notification has been sent.",
			},
			"log_drain_notification_sent": schema.BoolAttribute{
				Computed:    true,
				Description: "Indicates if a log drain notification has been sent.",
			},
			"port": schema.StringAttribute{
				Computed:    true,
				Description: "The port used by the server.",
			},
			"private_key_id": schema.Int64Attribute{
				Computed:    true,
				Description: "The ID of the private key associated with the server.",
			},
			"proxy": schema.SingleNestedAttribute{
				Computed:    true,
				Description: "Proxy settings for the server.",
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed:    true,
						Description: "The status of the proxy.",
					},
					"type": schema.StringAttribute{
						Computed:    true,
						Description: "The type of the proxy.",
					},
					"force_stop": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the proxy is forcefully stopped.",
					},
				},
			},
			"settings": schema.SingleNestedAttribute{
				Computed:    true,
				Description: "Settings related to the server.",
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Computed:    true,
						Description: "The ID of the server settings.",
					},
					"concurrent_builds": schema.Int64Attribute{
						Computed:    true,
						Description: "The number of concurrent builds allowed on the server.",
					},
					"delete_unused_networks": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if unused networks are deleted.",
					},
					"delete_unused_volumes": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if unused volumes are deleted.",
					},
					"docker_cleanup_frequency": schema.StringAttribute{
						Computed:    true,
						Description: "The frequency of Docker cleanup operations.",
					},
					"docker_cleanup_threshold": schema.Int64Attribute{
						Computed:    true,
						Description: "The threshold for Docker cleanup operations.",
					},
					"dynamic_timeout": schema.Int64Attribute{
						Computed:    true,
						Description: "The dynamic timeout setting for the server.",
					},
					"force_disabled": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the server is forcefully disabled.",
					},
					"force_docker_cleanup": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if Docker cleanup is forcefully executed.",
					},
					"generate_exact_labels": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if exact labels are generated.",
					},
					"is_build_server": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the server is a build server.",
					},
					"is_cloudflare_tunnel": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the server uses a Cloudflare tunnel.",
					},
					"is_jump_server": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the server is a jump server.",
					},
					"is_logdrain_axiom_enabled": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if log drain to Axiom is enabled.",
					},
					"is_logdrain_custom_enabled": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if custom log drain is enabled.",
					},
					"is_logdrain_highlight_enabled": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if log drain to Highlight is enabled.",
					},
					"is_logdrain_newrelic_enabled": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if log drain to New Relic is enabled.",
					},
					"is_metrics_enabled": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if metrics collection is enabled.",
					},
					"is_reachable": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the server is reachable.",
					},
					"is_server_api_enabled": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the server API is enabled.",
					},
					"is_swarm_manager": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the server is a Swarm manager.",
					},
					"is_swarm_worker": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the server is a Swarm worker.",
					},
					"is_usable": schema.BoolAttribute{
						Computed:    true,
						Description: "Indicates if the server is usable.",
					},
					"logdrain_axiom_api_key": schema.StringAttribute{
						Computed:    true,
						Description: "The API key for Axiom log drain.",
					},
					"logdrain_axiom_dataset_name": schema.StringAttribute{
						Computed:    true,
						Description: "The dataset name for Axiom log drain.",
					},
					"logdrain_custom_config": schema.StringAttribute{
						Computed:    true,
						Description: "The custom configuration for log drain.",
					},
					"logdrain_custom_config_parser": schema.StringAttribute{
						Computed:    true,
						Description: "The parser for custom log drain configuration.",
					},
					"logdrain_highlight_project_id": schema.StringAttribute{
						Computed:    true,
						Description: "The project ID for Highlight log drain.",
					},
					"logdrain_newrelic_base_uri": schema.StringAttribute{
						Computed:    true,
						Description: "The base URI for New Relic log drain.",
					},
					"logdrain_newrelic_license_key": schema.StringAttribute{
						Computed:    true,
						Description: "The license key for New Relic log drain.",
					},
					"metrics_history_days": schema.Int64Attribute{
						Computed:    true,
						Description: "The number of days to retain metrics history.",
					},
					"metrics_refresh_rate_seconds": schema.Int64Attribute{
						Computed:    true,
						Description: "The refresh rate for metrics in seconds.",
					},
					"metrics_token": schema.StringAttribute{
						Computed:    true,
						Description: "The token used for metrics collection.",
					},
					"server_id": schema.Int64Attribute{
						Computed:    true,
						Description: "The ID of the server.",
					},
					"server_timezone": schema.StringAttribute{
						Computed:    true,
						Description: "The timezone of the server.",
					},
					"wildcard_domain": schema.StringAttribute{
						Computed:    true,
						Description: "The wildcard domain associated with the server.",
					},
					"created_at": schema.StringAttribute{
						Computed:    true,
						Description: "The creation timestamp of the server.",
					},
					"updated_at": schema.StringAttribute{
						Computed:    true,
						Description: "The last update timestamp of the server.",
					},
				},
			},
			"swarm_cluster": schema.StringAttribute{
				Computed:    true,
				Description: "The swarm cluster associated with the server.",
			},
			"team_id": schema.Int64Attribute{
				Computed:    true,
				Description: "The ID of the team that owns the server.",
			},
			"unreachable_count": schema.Int64Attribute{
				Computed:    true,
				Description: "The count of times the server was unreachable.",
			},
			"unreachable_notification_sent": schema.BoolAttribute{
				Computed:    true,
				Description: "Indicates if an unreachable notification has been sent.",
			},
			"user": schema.StringAttribute{
				Computed:    true,
				Description: "The user associated with the server.",
			},
			"validation_logs": schema.StringAttribute{
				Computed:    true,
				Description: "Logs related to server validation.",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The creation timestamp of the server.",
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: "The last update timestamp of the server.",
			},
		},
	}
}

func (s *ServerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	resp.Diagnostics.Append(configure.ConfigureClient(ctx, req, &s.client)...)
}

func (s *ServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var server ServerModel

	diags := req.Config.Get(ctx, &server)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	serverSaved, diags := readServer(ctx, *s.client, server.UUID, server.Name)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	mapServerModel(&server, serverSaved)

	tflog.Trace(ctx, "Successfully read team data", map[string]interface{}{
		"server_uuid": serverSaved.UUID,
		"name":        serverSaved.Name,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &server)...)

}
