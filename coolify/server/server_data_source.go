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
				Optional: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"ip": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"high_disk_usage_notification_sent": schema.BoolAttribute{
				Computed: true,
			},
			"log_drain_notification_sent": schema.BoolAttribute{
				Computed: true,
			},
			"port": schema.StringAttribute{
				Computed: true,
			},
			"private_key_id": schema.Int64Attribute{
				Computed: true,
			},
			"proxy": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed: true,
					},
					"type": schema.StringAttribute{
						Computed: true,
					},
					"force_stop": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
			"settings": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Computed: true,
					},
					"concurrent_builds": schema.Int64Attribute{
						Computed: true,
					},
					"delete_unused_networks": schema.BoolAttribute{
						Computed: true,
					},
					"delete_unused_volumes": schema.BoolAttribute{
						Computed: true,
					},
					"docker_cleanup_frequency": schema.StringAttribute{
						Computed: true,
					},
					"docker_cleanup_threshold": schema.Int64Attribute{
						Computed: true,
					},
					"dynamic_timeout": schema.Int64Attribute{
						Computed: true,
					},
					"force_disabled": schema.BoolAttribute{
						Computed: true,
					},
					"force_docker_cleanup": schema.BoolAttribute{
						Computed: true,
					},
					"generate_exact_labels": schema.BoolAttribute{
						Computed: true,
					},
					"is_build_server": schema.BoolAttribute{
						Computed: true,
					},
					"is_cloudflare_tunnel": schema.BoolAttribute{
						Computed: true,
					},
					"is_jump_server": schema.BoolAttribute{
						Computed: true,
					},
					"is_logdrain_axiom_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"is_logdrain_custom_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"is_logdrain_highlight_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"is_logdrain_newrelic_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"is_metrics_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"is_reachable": schema.BoolAttribute{
						Computed: true,
					},
					"is_server_api_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"is_swarm_manager": schema.BoolAttribute{
						Computed: true,
					},
					"is_swarm_worker": schema.BoolAttribute{
						Computed: true,
					},
					"is_usable": schema.BoolAttribute{
						Computed: true,
					},
					"logdrain_axiom_api_key": schema.StringAttribute{
						Computed: true,
					},
					"logdrain_axiom_dataset_name": schema.StringAttribute{
						Computed: true,
					},
					"logdrain_custom_config": schema.StringAttribute{
						Computed: true,
					},
					"logdrain_custom_config_parser": schema.StringAttribute{
						Computed: true,
					},
					"logdrain_highlight_project_id": schema.StringAttribute{
						Computed: true,
					},
					"logdrain_newrelic_base_uri": schema.StringAttribute{
						Computed: true,
					},
					"logdrain_newrelic_license_key": schema.StringAttribute{
						Computed: true,
					},
					"metrics_history_days": schema.Int64Attribute{
						Computed: true,
					},
					"metrics_refresh_rate_seconds": schema.Int64Attribute{
						Computed: true,
					},
					"metrics_token": schema.StringAttribute{
						Computed: true,
					},
					"server_id": schema.Int64Attribute{
						Computed: true,
					},
					"server_timezone": schema.StringAttribute{
						Computed: true,
					},
					"wildcard_domain": schema.StringAttribute{
						Computed: true,
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"swarm_cluster": schema.StringAttribute{
				Computed: true,
			},
			"team_id": schema.Int64Attribute{
				Computed: true,
			},
			"unreachable_count": schema.Int64Attribute{
				Computed: true,
			},
			"unreachable_notification_sent": schema.BoolAttribute{
				Computed: true,
			},
			"user": schema.StringAttribute{
				Computed: true,
			},
			"validation_logs": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
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
