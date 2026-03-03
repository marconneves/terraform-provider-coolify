package server

import (
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
	coolify_sdk "github.com/marconneves/coolify-sdk-go"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

type ServerModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Description    types.String `tfsdk:"description"`
	IP             types.String `tfsdk:"ip"`
	Port           types.Int32  `tfsdk:"port"`
	User           types.String `tfsdk:"user"`
	PrivateKeyUUID types.String `tfsdk:"private_key_uuid"`
}

type ServerDataSourceModel struct {
	ServerModel
	HighDiskUsageNotificationSent types.Bool     `tfsdk:"high_disk_usage_notification_sent"`
	LogDrainNotificationSent      types.Bool     `tfsdk:"log_drain_notification_sent"`
	Proxy                         *ProxyModel    `tfsdk:"proxy"`
	Settings                      *SettingsModel `tfsdk:"settings"`
	SwarmCluster                  types.String   `tfsdk:"swarm_cluster"`
	TeamID                        types.Int64    `tfsdk:"team_id"`
	UnreachableCount              types.Int64    `tfsdk:"unreachable_count"`
	UnreachableNotificationSent   types.Bool     `tfsdk:"unreachable_notification_sent"`
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

func mapCommonServerFields(data *ServerModel, server *coolify_sdk.Server) {
	data.ID = types.StringValue(server.UUID)
	data.IP = types.StringValue(server.IP)
	data.Name = types.StringValue(server.Name)
	data.Description = configure.ValueStringValue(server.Description, data.Description)
	data.User = types.StringValue(server.User)

	data.Port = types.Int32Value(int32(server.Port))
}

func mapServerDataSourceModel(data *ServerDataSourceModel, server *coolify_sdk.Server) {
	mapCommonServerFields(&data.ServerModel, server)

	// TODO: Get private key uuid founding in api
	data.PrivateKeyUUID = types.StringValue(server.UUID)

	data.HighDiskUsageNotificationSent = types.BoolValue(server.HighDiskUsageNotificationSent)
	data.LogDrainNotificationSent = types.BoolValue(server.LogDrainNotificationSent)

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
		settings.LogdrainAxiomApiKey = configure.ValueStringValue(server.Settings.LogdrainAxiomApiKey, settings.LogdrainAxiomApiKey)
		settings.LogdrainAxiomDatasetName = configure.ValueStringValue(server.Settings.LogdrainAxiomDatasetName, settings.LogdrainAxiomDatasetName)
		settings.LogdrainCustomConfig = configure.ValueStringValue(server.Settings.LogdrainCustomConfig, settings.LogdrainCustomConfig)
		settings.LogdrainCustomConfigParser = configure.ValueStringValue(server.Settings.LogdrainCustomConfigParser, settings.LogdrainCustomConfigParser)
		settings.LogdrainHighlightProjectId = configure.ValueStringValue(server.Settings.LogdrainHighlightProjectId, settings.LogdrainHighlightProjectId)
		settings.LogdrainNewRelicBaseUri = configure.ValueStringValue(server.Settings.LogdrainNewRelicBaseUri, settings.LogdrainNewRelicBaseUri)
		settings.LogdrainNewRelicLicenseKey = configure.ValueStringValue(server.Settings.LogdrainNewRelicLicenseKey, settings.LogdrainNewRelicLicenseKey)
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

	data.SwarmCluster = configure.ValueStringValue(server.SwarmCluster, data.SwarmCluster)
	data.TeamID = types.Int64Value(int64(server.TeamID))
	data.UnreachableCount = types.Int64Value(int64(server.UnreachableCount))
	data.UnreachableNotificationSent = types.BoolValue(server.UnreachableNotificationSent)
	data.ValidationLogs = configure.ValueStringValue(server.ValidationLogs, data.ValidationLogs)

	data.CreatedAt = types.StringValue(server.CreatedAt.Format(time.RFC3339))
	data.UpdatedAt = types.StringValue(server.UpdatedAt.Format(time.RFC3339))
}

func mapServerResourceModel(projectData *ServerModel, project *coolify_sdk.Server) {
	mapCommonServerFields(projectData, project)
}
