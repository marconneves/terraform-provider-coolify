package application_dockerimage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/marconneves/coolify-sdk-go/application"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// ApplicationDockerImageModel describes the Terraform resource state for an
// application backed by a prebuilt Docker image (e.g. GCP Artifact Registry).
type ApplicationDockerImageModel struct {
	Id              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	ProjectUUID     types.String `tfsdk:"project_uuid"`
	ServerUUID      types.String `tfsdk:"server_uuid"`
	EnvironmentName types.String `tfsdk:"environment_name"`
	EnvironmentUUID types.String `tfsdk:"environment_uuid"`
	DestinationUUID types.String `tfsdk:"destination_uuid"`

	DockerRegistryImageName types.String `tfsdk:"docker_registry_image_name"`
	DockerRegistryImageTag  types.String `tfsdk:"docker_registry_image_tag"`

	PortsExposes  types.String `tfsdk:"ports_exposes"`
	PortsMappings types.String `tfsdk:"ports_mappings"`
	Domains       types.String `tfsdk:"domains"`

	InstantDeploy          types.Bool `tfsdk:"instant_deploy"`
	IsForceHTTPSEnabled    types.Bool `tfsdk:"is_force_https_enabled"`
	ConnectToDockerNetwork types.Bool `tfsdk:"connect_to_docker_network"`
	RedeployOnUpdate       types.Bool `tfsdk:"redeploy_on_update"`

	HealthCheckEnabled      types.Bool   `tfsdk:"health_check_enabled"`
	HealthCheckPath         types.String `tfsdk:"health_check_path"`
	HealthCheckPort         types.String `tfsdk:"health_check_port"`
	HealthCheckHost         types.String `tfsdk:"health_check_host"`
	HealthCheckMethod       types.String `tfsdk:"health_check_method"`
	HealthCheckReturnCode   types.Int64  `tfsdk:"health_check_return_code"`
	HealthCheckScheme       types.String `tfsdk:"health_check_scheme"`
	HealthCheckResponseText types.String `tfsdk:"health_check_response_text"`
	HealthCheckInterval     types.Int64  `tfsdk:"health_check_interval"`
	HealthCheckTimeout      types.Int64  `tfsdk:"health_check_timeout"`
	HealthCheckRetries      types.Int64  `tfsdk:"health_check_retries"`
	HealthCheckStartPeriod  types.Int64  `tfsdk:"health_check_start_period"`

	LimitsMemory            types.String `tfsdk:"limits_memory"`
	LimitsMemorySwap        types.String `tfsdk:"limits_memory_swap"`
	LimitsMemorySwappiness  types.Int64  `tfsdk:"limits_memory_swappiness"`
	LimitsMemoryReservation types.String `tfsdk:"limits_memory_reservation"`
	LimitsCpus              types.String `tfsdk:"limits_cpus"`
	LimitsCpuset            types.String `tfsdk:"limits_cpuset"`
	LimitsCPUShares         types.Int64  `tfsdk:"limits_cpu_shares"`

	CustomLabels           types.String `tfsdk:"custom_labels"`
	CustomDockerRunOptions types.String `tfsdk:"custom_docker_run_options"`

	EnvironmentVariables types.Map `tfsdk:"environment_variables"`

	Status types.String `tfsdk:"status"`
	FQDN   types.String `tfsdk:"fqdn"`
}

// mapApplicationToModel updates the model with the application returned by the
// API, preserving the user's plan values for write-only/unknown fields.
func mapApplicationToModel(data *ApplicationDockerImageModel, app *application.Application) {
	data.Id = types.StringValue(app.UUID)
	data.Name = types.StringValue(app.Name)
	data.Description = configure.ValueStringValue(app.Description, data.Description)
	data.DockerRegistryImageName = configure.ValueStringValue(app.DockerRegistryImageName, data.DockerRegistryImageName)
	data.DockerRegistryImageTag = configure.ValueStringValue(app.DockerRegistryImageTag, data.DockerRegistryImageTag)
	data.PortsExposes = types.StringValue(app.PortsExposes)
	data.PortsMappings = configure.ValueStringValue(app.PortsMappings, data.PortsMappings)
	data.Domains = configure.ValueStringValue(app.FQDN, data.Domains)

	data.HealthCheckEnabled = types.BoolValue(app.HealthCheckEnabled)
	data.HealthCheckPath = types.StringValue(app.HealthCheckPath)
	data.HealthCheckPort = configure.ValueStringValue(app.HealthCheckPort, data.HealthCheckPort)
	data.HealthCheckHost = configure.ValueStringValue(app.HealthCheckHost, data.HealthCheckHost)
	data.HealthCheckMethod = types.StringValue(app.HealthCheckMethod)
	data.HealthCheckReturnCode = types.Int64Value(int64(app.HealthCheckReturnCode))
	data.HealthCheckScheme = types.StringValue(app.HealthCheckScheme)
	data.HealthCheckResponseText = configure.ValueStringValue(app.HealthCheckResponseText, data.HealthCheckResponseText)
	data.HealthCheckInterval = types.Int64Value(int64(app.HealthCheckInterval))
	data.HealthCheckTimeout = types.Int64Value(int64(app.HealthCheckTimeout))
	data.HealthCheckRetries = types.Int64Value(int64(app.HealthCheckRetries))
	data.HealthCheckStartPeriod = types.Int64Value(int64(app.HealthCheckStartPeriod))

	data.LimitsMemory = types.StringValue(app.LimitsMemory)
	data.LimitsMemorySwap = types.StringValue(app.LimitsMemorySwap)
	data.LimitsMemorySwappiness = types.Int64Value(int64(app.LimitsMemorySwappiness))
	data.LimitsMemoryReservation = types.StringValue(app.LimitsMemoryReservation)
	data.LimitsCpus = types.StringValue(app.LimitsCpus)
	data.LimitsCpuset = configure.ValueStringValue(app.LimitsCpuset, data.LimitsCpuset)
	data.LimitsCPUShares = types.Int64Value(int64(app.LimitsCPUShares))

	data.CustomLabels = configure.ValueStringValue(app.CustomLabels, data.CustomLabels)
	data.CustomDockerRunOptions = configure.ValueStringValue(app.CustomDockerRunOptions, data.CustomDockerRunOptions)

	data.Status = types.StringValue(app.Status)
	data.FQDN = configure.ValueStringValue(app.FQDN, types.StringNull())
}

// mapEnvsToModel updates the EnvironmentVariables map on the model from the
// envs returned by the API. Only keys present in the user's plan are mirrored
// back to avoid surfacing values managed by other resources (e.g.
// coolify_application_env or auto-generated by Coolify).
func mapEnvsToModel(ctx context.Context, data *ApplicationDockerImageModel, envs []application.EnvironmentVariable) diag.Diagnostics {
	tracked := map[string]bool{}
	if !data.EnvironmentVariables.IsNull() && !data.EnvironmentVariables.IsUnknown() {
		for k := range data.EnvironmentVariables.Elements() {
			tracked[k] = true
		}
	}

	values := map[string]string{}
	for _, e := range envs {
		if !tracked[e.Key] {
			continue
		}
		values[e.Key] = e.Value
	}

	if len(values) == 0 && (data.EnvironmentVariables.IsNull() || data.EnvironmentVariables.IsUnknown()) {
		data.EnvironmentVariables = types.MapNull(types.StringType)
		return nil
	}

	m, diags := types.MapValueFrom(ctx, types.StringType, values)
	if diags.HasError() {
		return diags
	}
	data.EnvironmentVariables = m
	return nil
}
