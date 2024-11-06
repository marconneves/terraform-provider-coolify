---
subcategory: "Server"
page_title: "Coolify: coolify_server"
description: |-
  The `coolify_server` data source allows you to retrieve detailed information about a server by its UUID or name.
---

# Data Source: `coolify_server`

The `coolify_server` data source enables you to fetch comprehensive details about a specific server in Coolify by specifying either its `uuid` or `name`. This includes network information, proxy settings, server configurations, and monitoring data.

## Example Usage

```hcl
data "coolify_server" "example" {
  uuid = "server-uuid-1234"
}

output "server_uuid" {
  value = data.coolify_server.example.uuid
}

output "server_name" {
  value = data.coolify_server.example.name
}

output "server_ip" {
  value = data.coolify_server.example.ip
}

output "proxy_status" {
  value = data.coolify_server.example.proxy.status
}

output "docker_cleanup_frequency" {
  value = data.coolify_server.example.settings.docker_cleanup_frequency
}
```

```hcl
data "coolify_server" "example_by_name" {
  name = "example-server"
}

output "server_description" {
  value = data.coolify_server.example_by_name.description
}

output "metrics_token" {
  value = data.coolify_server.example_by_name.settings.metrics_token
}
```

## Input Parameters

- `uuid` (Optional): The unique identifier of the server. If specified, the data source will retrieve the server data associated with this UUID.
- `name` (Optional): The name of the server. If specified and `uuid` is not provided, the data source will retrieve the server data associated with this name.

**Note**: At least one of `uuid` or `name` must be specified to identify the server.

## Output Attributes

### Main Attributes

- `uuid` (String): The unique identifier of the server.
- `name` (String): The name of the server.
- `ip` (String): The IP address of the server.
- `description` (String): A description of the server.
- `high_disk_usage_notification_sent` (Bool): Indicates whether a high disk usage notification has been sent.
- `log_drain_notification_sent` (Bool): Indicates whether a log drain notification has been sent.
- `port` (String): The port on which the server is running.
- `private_key_id` (Int64): The identifier of the private key associated with the server.
- `swarm_cluster` (String): The Swarm cluster to which the server belongs.
- `team_id` (Int64): The ID of the team associated with the server.
- `unreachable_count` (Int64): The number of times the server has been marked as unreachable.
- `unreachable_notification_sent` (Bool): Indicates whether an unreachable notification has been sent.
- `user` (String): The user associated with the server.
- `validation_logs` (String): Logs related to the server's validation processes.
- `created_at` (String): The timestamp when the server was created, in RFC3339 format.
- `updated_at` (String): The timestamp when the server was last updated, in RFC3339 format.

### Proxy

Represents the proxy configuration of the server.

- `proxy` (Object):
  - `status` (String): The current status of the proxy.
  - `type` (String): The type of proxy being used.
  - `force_stop` (Bool): Indicates whether the proxy has been forcefully stopped.

### Settings

Contains various configuration settings related to the server's operations, including Docker management, monitoring, and logging configurations.

- `settings` (Object):
  - `id` (Int64): The identifier for the settings.
  - `concurrent_builds` (Int64): The number of concurrent builds allowed.
  - `delete_unused_networks` (Bool): Whether to delete unused networks.
  - `delete_unused_volumes` (Bool): Whether to delete unused volumes.
  - `docker_cleanup_frequency` (String): The frequency at which Docker cleanup occurs.
  - `docker_cleanup_threshold` (Int64): The threshold for Docker cleanup.
  - `dynamic_timeout` (Int64): The dynamic timeout setting for the server.
  - `force_disabled` (Bool): Whether the server has been forcefully disabled.
  - `force_docker_cleanup` (Bool): Whether Docker cleanup is forced.
  - `generate_exact_labels` (Bool): Whether to generate exact labels.
  - `is_build_server` (Bool): Indicates if the server is used for builds.
  - `is_cloudflare_tunnel` (Bool): Indicates if Cloudflare tunnel is enabled.
  - `is_jump_server` (Bool): Indicates if the server is a jump server.
  - `is_logdrain_axiom_enabled` (Bool): Indicates if Axiom logdrain is enabled.
  - `is_logdrain_custom_enabled` (Bool): Indicates if custom logdrain is enabled.
  - `is_logdrain_highlight_enabled` (Bool): Indicates if Highlight logdrain is enabled.
  - `is_logdrain_newrelic_enabled` (Bool): Indicates if New Relic logdrain is enabled.
  - `is_metrics_enabled` (Bool): Indicates if metrics collection is enabled.
  - `is_reachable` (Bool): Indicates if the server is reachable.
  - `is_server_api_enabled` (Bool): Indicates if the server API is enabled.
  - `is_swarm_manager` (Bool): Indicates if the server is a Swarm manager.
  - `is_swarm_worker` (Bool): Indicates if the server is a Swarm worker.
  - `is_usable` (Bool): Indicates if the server is usable.
  - `logdrain_axiom_api_key` (String): The API key for Axiom logdrain.
  - `logdrain_axiom_dataset_name` (String): The dataset name for Axiom logdrain.
  - `logdrain_custom_config` (String): The custom configuration for logdrain.
  - `logdrain_custom_config_parser` (String): The custom config parser for logdrain.
  - `logdrain_highlight_project_id` (String): The project ID for Highlight logdrain.
  - `logdrain_newrelic_base_uri` (String): The base URI for New Relic logdrain.
  - `logdrain_newrelic_license_key` (String): The license key for New Relic logdrain.
  - `metrics_history_days` (Int64): The number of days to retain metrics history.
  - `metrics_refresh_rate_seconds` (Int64): The refresh rate for metrics in seconds.
  - `metrics_token` (String): The token used for metrics collection.
  - `server_id` (Int64): The server identifier used for configuration purposes.
  - `server_timezone` (String): The timezone of the server.
  - `wildcard_domain` (String): The wildcard domain associated with the server.
  - `created_at` (String): The timestamp when the settings were created, in RFC3339 format.
  - `updated_at` (String): The timestamp when the settings were last updated, in RFC3339 format.