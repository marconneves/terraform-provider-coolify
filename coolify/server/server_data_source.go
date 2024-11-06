package server

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

type ServerDataSourceModel struct {
	UUID                          types.String   `tfsdk:"uuid"`
	Name                          types.String   `tfsdk:"name"`
	IP                            types.String   `tfsdk:"ip"`
	Description                   types.String   `tfsdk:"description"`
	HighDiskUsageNotificationSent types.Bool     `tfsdk:"high_disk_usage_notification_sent"`
	LogDrainNotificationSent      types.Bool     `tfsdk:"log_drain_notification_sent"`
	Port                          types.String   `tfsdk:"port"`
	PrivateKeyID                  types.Int64    `tfsdk:"private_key_id"`
	Proxy                         *ProxyModel    `tfsdk:"proxy"`
	Settings                      *SettingsModel `tfsdk:"settings"`
	SwarmCluster                  types.String   `tfsdk:"swarm_cluster"`
	TeamID                        types.Int64    `tfsdk:"team_id"`
	UnreachableCount              types.Int64    `tfsdk:"unreachable_count"`
	UnreachableNotificationSent   types.Bool     `tfsdk:"unreachable_notification_sent"`
	User                          types.String   `tfsdk:"user"`
	ValidationLogs                types.String   `tfsdk:"validation_logs"`
	CreatedAt                     types.String   `tfsdk:"created_at"`
	UpdatedAt                     types.String   `tfsdk:"updated_at"`
}

type ProxyModel struct {
	Status    types.String `tfsdk:"status"`
	Type      types.String `tfsdk:"type"`
	ForceStop types.Bool   `tfsdk:"force_stop"`
}

type SettingsModel struct {
	Id                         types.Int64  `tfsdk:"id"`
	ConcurrentBuilds           types.Int64  `tfsdk:"concurrent_builds"`
	DeleteUnusedNetworks       types.Bool   `tfsdk:"delete_unused_networks"`
	DeleteUnusedVolumes        types.Bool   `tfsdk:"delete_unused_volumes"`
	DockerCleanupFrequency     types.String `tfsdk:"docker_cleanup_frequency"`
	DockerCleanupThreshold     types.Int64  `tfsdk:"docker_cleanup_threshold"`
	DynamicTimeout             types.Int64  `tfsdk:"dynamic_timeout"`
	ForceDisabled              types.Bool   `tfsdk:"force_disabled"`
	ForceDockerCleanup         types.Bool   `tfsdk:"force_docker_cleanup"`
	GenerateExactLabels        types.Bool   `tfsdk:"generate_exact_labels"`
	IsBuildServer              types.Bool   `tfsdk:"is_build_server"`
	IsCloudflareTunnel         types.Bool   `tfsdk:"is_cloudflare_tunnel"`
	IsJumpServer               types.Bool   `tfsdk:"is_jump_server"`
	IsLogdrainAxiomEnabled     types.Bool   `tfsdk:"is_logdrain_axiom_enabled"`
	IsLogdrainCustomEnabled    types.Bool   `tfsdk:"is_logdrain_custom_enabled"`
	IsLogdrainHighlightEnabled types.Bool   `tfsdk:"is_logdrain_highlight_enabled"`
	IsLogdrainNewRelicEnabled  types.Bool   `tfsdk:"is_logdrain_newrelic_enabled"`
	IsMetricsEnabled           types.Bool   `tfsdk:"is_metrics_enabled"`
	IsReachable                types.Bool   `tfsdk:"is_reachable"`
	IsServerAPIEnabled         types.Bool   `tfsdk:"is_server_api_enabled"`
	IsSwarmManager             types.Bool   `tfsdk:"is_swarm_manager"`
	IsSwarmWorker              types.Bool   `tfsdk:"is_swarm_worker"`
	IsUsable                   types.Bool   `tfsdk:"is_usable"`
	LogdrainAxiomApiKey        types.String `tfsdk:"logdrain_axiom_api_key"`
	LogdrainAxiomDatasetName   types.String `tfsdk:"logdrain_axiom_dataset_name"`
	LogdrainCustomConfig       types.String `tfsdk:"logdrain_custom_config"`
	LogdrainCustomConfigParser types.String `tfsdk:"logdrain_custom_config_parser"`
	LogdrainHighlightProjectId types.String `tfsdk:"logdrain_highlight_project_id"`
	LogdrainNewRelicBaseUri    types.String `tfsdk:"logdrain_newrelic_base_uri"`
	LogdrainNewRelicLicenseKey types.String `tfsdk:"logdrain_newrelic_license_key"`
	MetricsHistoryDays         types.Int64  `tfsdk:"metrics_history_days"`
	MetricsRefreshRateSeconds  types.Int64  `tfsdk:"metrics_refresh_rate_seconds"`
	MetricsToken               types.String `tfsdk:"metrics_token"`
	ServerId                   types.Int64  `tfsdk:"server_id"`
	ServerTimezone             types.String `tfsdk:"server_timezone"`
	WildcardDomain             types.String `tfsdk:"wildcard_domain"`
	CreatedAt                  types.String `tfsdk:"created_at"`
	UpdatedAt                  types.String `tfsdk:"updated_at"`
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
	var data ServerDataSourceModel

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var server *coolify_sdk.Server
	var err error

	if !data.UUID.IsNull() {
		server, err = s.client.Server.Get(data.UUID.ValueString())

		if err != nil {
			resp.Diagnostics.AddError(
				"Error Reading Server",
				"Could not read server UUID "+data.UUID.ValueString()+": "+err.Error(),
			)
			return
		}
	} else if !data.Name.IsNull() {
		servers, err := s.client.Server.List()

		if err != nil {
			resp.Diagnostics.AddError(
				"Error Listing Servers",
				"Could not list servers: "+err.Error(),
			)
			return
		}

		var serverId *string

		for _, serverItem := range *servers {
			if serverItem.Name == data.Name.ValueString() {
				serverId = &serverItem.UUID
				break
			}
		}

		if serverId == nil {
			resp.Diagnostics.AddError(
				"Server Not Found",
				"Could not find server with name "+data.Name.ValueString(),
			)
			return
		}

		server, err = s.client.Server.Get(*serverId)

		if err != nil {
			resp.Diagnostics.AddError(
				"Server Not Found",
				"Could not find server with name "+data.Name.ValueString(),
			)
			return
		}

	} else {
		resp.Diagnostics.AddError(
			"Missing Configuration",
			"Either 'id' or 'name' must be specified",
		)
		return
	}

	data.UUID = types.StringValue(server.UUID)
	data.IP = types.StringValue(server.IP)
	data.Name = types.StringValue(server.Name)
	data.Description = types.StringNull()
	if server.Description != nil {
		data.Description = types.StringValue(*server.Description)
	}

	data.HighDiskUsageNotificationSent = types.BoolValue(server.HighDiskUsageNotificationSent)
	data.LogDrainNotificationSent = types.BoolValue(server.LogDrainNotificationSent)
	data.Port = types.StringValue(server.Port)
	data.PrivateKeyID = types.Int64Value(int64(server.PrivateKeyID))

	if server.Proxy != nil {
		proxy := ProxyModel{}

		proxy.Status = types.StringValue(server.Proxy.Status)
		proxy.Type = types.StringValue(server.Proxy.Type)
		proxy.ForceStop = types.BoolValue(server.Proxy.ForceStop)

		data.Proxy = &proxy
	}

	if server.Settings != nil {
		settings := SettingsModel{}

		settings.Id = types.Int64Value(int64(server.Settings.Id))
		settings.ConcurrentBuilds = types.Int64Value(int64(server.Settings.ConcurrentBuilds))
		settings.DeleteUnusedNetworks = types.BoolValue(server.Settings.DeleteUnusedNetworks)
		settings.DeleteUnusedVolumes = types.BoolValue(server.Settings.DeleteUnusedVolumes)
		settings.DockerCleanupFrequency = types.StringValue(server.Settings.DockerCleanupFrequency)
		settings.DockerCleanupThreshold = types.Int64Value(int64(server.Settings.DockerCleanupThreshold))
		settings.DynamicTimeout = types.Int64Value(int64(server.Settings.DynamicTimeout))
		settings.ForceDisabled = types.BoolValue(server.Settings.ForceDisabled)
		settings.ForceDockerCleanup = types.BoolValue(server.Settings.ForceDockerCleanup)
		settings.GenerateExactLabels = types.BoolValue(server.Settings.GenerateExactLabels)
		settings.IsBuildServer = types.BoolValue(server.Settings.IsBuildServer)
		settings.IsCloudflareTunnel = types.BoolValue(server.Settings.IsCloudflareTunnel)
		settings.IsJumpServer = types.BoolValue(server.Settings.IsJumpServer)
		settings.IsLogdrainAxiomEnabled = types.BoolValue(server.Settings.IsLogdrainAxiomEnabled)
		settings.IsLogdrainCustomEnabled = types.BoolValue(server.Settings.IsLogdrainCustomEnabled)
		settings.IsLogdrainHighlightEnabled = types.BoolValue(server.Settings.IsLogdrainHighlightEnabled)
		settings.IsLogdrainNewRelicEnabled = types.BoolValue(server.Settings.IsLogdrainNewRelicEnabled)
		settings.IsMetricsEnabled = types.BoolValue(server.Settings.IsMetricsEnabled)
		settings.IsReachable = types.BoolValue(server.Settings.IsReachable)
		settings.IsServerAPIEnabled = types.BoolValue(server.Settings.IsServerAPIEnabled)
		settings.IsSwarmManager = types.BoolValue(server.Settings.IsSwarmManager)
		settings.IsSwarmWorker = types.BoolValue(server.Settings.IsSwarmWorker)
		settings.IsUsable = types.BoolValue(server.Settings.IsUsable)
		settings.LogdrainAxiomApiKey = types.StringNull()
		if server.Settings.LogdrainAxiomApiKey != nil {
			settings.LogdrainAxiomApiKey = types.StringValue(*server.Settings.LogdrainAxiomApiKey)
		}
		settings.LogdrainAxiomDatasetName = types.StringNull()
		if server.Settings.LogdrainAxiomDatasetName != nil {
			settings.LogdrainAxiomDatasetName = types.StringValue(*server.Settings.LogdrainAxiomDatasetName)
		}
		settings.LogdrainCustomConfig = types.StringNull()
		if server.Settings.LogdrainCustomConfig != nil {
			settings.LogdrainCustomConfig = types.StringValue(*server.Settings.LogdrainCustomConfig)
		}
		settings.LogdrainCustomConfigParser = types.StringNull()
		if server.Settings.LogdrainCustomConfigParser != nil {
			settings.LogdrainCustomConfigParser = types.StringValue(*server.Settings.LogdrainCustomConfigParser)
		}
		settings.LogdrainHighlightProjectId = types.StringNull()
		if server.Settings.LogdrainHighlightProjectId != nil {
			settings.LogdrainHighlightProjectId = types.StringValue(*server.Settings.LogdrainHighlightProjectId)
		}
		settings.LogdrainNewRelicBaseUri = types.StringNull()
		if server.Settings.LogdrainNewRelicBaseUri != nil {
			settings.LogdrainNewRelicBaseUri = types.StringValue(*server.Settings.LogdrainNewRelicBaseUri)
		}
		settings.LogdrainNewRelicLicenseKey = types.StringNull()
		if server.Settings.LogdrainNewRelicLicenseKey != nil {
			settings.LogdrainNewRelicLicenseKey = types.StringValue(*server.Settings.LogdrainNewRelicLicenseKey)
		}
		settings.MetricsHistoryDays = types.Int64Value(int64(server.Settings.MetricsHistoryDays))
		settings.MetricsRefreshRateSeconds = types.Int64Value(int64(server.Settings.MetricsRefreshRateSeconds))
		settings.MetricsToken = types.StringValue(server.Settings.MetricsToken)
		settings.ServerId = types.Int64Value(int64(server.Settings.ServerId))
		settings.ServerTimezone = types.StringValue(server.Settings.ServerTimezone)
		settings.WildcardDomain = types.StringNull()
		if server.Settings.WildcardDomain != nil {
			settings.WildcardDomain = types.StringValue(*server.Settings.WildcardDomain)
		}

		settings.CreatedAt = types.StringValue(server.Settings.CreatedAt.Format(time.RFC3339))
		settings.UpdatedAt = types.StringValue(server.Settings.UpdatedAt.Format(time.RFC3339))

		data.Settings = &settings

	}

	data.SwarmCluster = types.StringNull()
	if server.SwarmCluster != nil {
		data.SwarmCluster = types.StringValue(*server.SwarmCluster)
	}
	data.TeamID = types.Int64Value(int64(server.TeamID))
	data.UnreachableCount = types.Int64Value(int64(server.UnreachableCount))
	data.UnreachableNotificationSent = types.BoolValue(server.UnreachableNotificationSent)
	data.User = types.StringValue(server.User)
	data.ValidationLogs = types.StringNull()
	if server.ValidationLogs != nil {
		data.ValidationLogs = types.StringValue(*server.ValidationLogs)
	}

	data.CreatedAt = types.StringValue(server.CreatedAt.Format(time.RFC3339))
	data.UpdatedAt = types.StringValue(server.UpdatedAt.Format(time.RFC3339))

	diags = resp.State.Set(ctx, &data)

	tflog.Trace(ctx, "read a team by ID or Name data source")

	resp.Diagnostics.Append(diags...)

}
