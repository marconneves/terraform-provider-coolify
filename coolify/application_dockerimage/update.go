package application_dockerimage

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/application"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// runtimeFieldsChanged reports whether any field that requires the
// container to be redeployed differs between plan and state. Changes to
// metadata-only fields (name, description) and to the provider-side
// redeploy_on_update toggle itself do not count.
func runtimeFieldsChanged(plan, state ApplicationDockerImageModel) bool {
	return !plan.DockerRegistryImageName.Equal(state.DockerRegistryImageName) ||
		!plan.DockerRegistryImageTag.Equal(state.DockerRegistryImageTag) ||
		!plan.EnvironmentVariables.Equal(state.EnvironmentVariables) ||
		!plan.PortsExposes.Equal(state.PortsExposes) ||
		!plan.PortsMappings.Equal(state.PortsMappings) ||
		!plan.Domains.Equal(state.Domains) ||
		!plan.IsForceHTTPSEnabled.Equal(state.IsForceHTTPSEnabled) ||
		!plan.ConnectToDockerNetwork.Equal(state.ConnectToDockerNetwork) ||
		!plan.HealthCheckEnabled.Equal(state.HealthCheckEnabled) ||
		!plan.HealthCheckPath.Equal(state.HealthCheckPath) ||
		!plan.HealthCheckPort.Equal(state.HealthCheckPort) ||
		!plan.HealthCheckHost.Equal(state.HealthCheckHost) ||
		!plan.HealthCheckMethod.Equal(state.HealthCheckMethod) ||
		!plan.HealthCheckReturnCode.Equal(state.HealthCheckReturnCode) ||
		!plan.HealthCheckScheme.Equal(state.HealthCheckScheme) ||
		!plan.HealthCheckResponseText.Equal(state.HealthCheckResponseText) ||
		!plan.HealthCheckInterval.Equal(state.HealthCheckInterval) ||
		!plan.HealthCheckTimeout.Equal(state.HealthCheckTimeout) ||
		!plan.HealthCheckRetries.Equal(state.HealthCheckRetries) ||
		!plan.HealthCheckStartPeriod.Equal(state.HealthCheckStartPeriod) ||
		!plan.LimitsMemory.Equal(state.LimitsMemory) ||
		!plan.LimitsMemorySwap.Equal(state.LimitsMemorySwap) ||
		!plan.LimitsMemorySwappiness.Equal(state.LimitsMemorySwappiness) ||
		!plan.LimitsMemoryReservation.Equal(state.LimitsMemoryReservation) ||
		!plan.LimitsCpus.Equal(state.LimitsCpus) ||
		!plan.LimitsCpuset.Equal(state.LimitsCpuset) ||
		!plan.LimitsCPUShares.Equal(state.LimitsCPUShares) ||
		!plan.CustomLabels.Equal(state.CustomLabels) ||
		!plan.CustomDockerRunOptions.Equal(state.CustomDockerRunOptions)
}

func (r *ApplicationDockerImageResource) UpdateApplication(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ApplicationDockerImageModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ApplicationDockerImageModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateDTO := application.UpdateApplicationDTO{
		Name:                    configure.DiffString(plan.Name, state.Name),
		Description:             configure.DiffString(plan.Description, state.Description),
		DockerRegistryImageName: configure.DiffString(plan.DockerRegistryImageName, state.DockerRegistryImageName),
		DockerRegistryImageTag:  configure.DiffString(plan.DockerRegistryImageTag, state.DockerRegistryImageTag),
		PortsExposes:            configure.DiffString(plan.PortsExposes, state.PortsExposes),
		PortsMappings:           configure.DiffString(plan.PortsMappings, state.PortsMappings),
		Domains:                 configure.DiffString(plan.Domains, state.Domains),

		HealthCheckEnabled:      configure.DiffBool(plan.HealthCheckEnabled, state.HealthCheckEnabled),
		HealthCheckPath:         configure.DiffString(plan.HealthCheckPath, state.HealthCheckPath),
		HealthCheckPort:         configure.DiffString(plan.HealthCheckPort, state.HealthCheckPort),
		HealthCheckHost:         configure.DiffString(plan.HealthCheckHost, state.HealthCheckHost),
		HealthCheckMethod:       configure.DiffString(plan.HealthCheckMethod, state.HealthCheckMethod),
		HealthCheckReturnCode:   configure.DiffInt64(plan.HealthCheckReturnCode, state.HealthCheckReturnCode),
		HealthCheckScheme:       configure.DiffString(plan.HealthCheckScheme, state.HealthCheckScheme),
		HealthCheckResponseText: configure.DiffString(plan.HealthCheckResponseText, state.HealthCheckResponseText),
		HealthCheckInterval:     configure.DiffInt64(plan.HealthCheckInterval, state.HealthCheckInterval),
		HealthCheckTimeout:      configure.DiffInt64(plan.HealthCheckTimeout, state.HealthCheckTimeout),
		HealthCheckRetries:      configure.DiffInt64(plan.HealthCheckRetries, state.HealthCheckRetries),
		HealthCheckStartPeriod:  configure.DiffInt64(plan.HealthCheckStartPeriod, state.HealthCheckStartPeriod),

		LimitsMemory:            configure.DiffString(plan.LimitsMemory, state.LimitsMemory),
		LimitsMemorySwap:        configure.DiffString(plan.LimitsMemorySwap, state.LimitsMemorySwap),
		LimitsMemorySwappiness:  configure.DiffInt64(plan.LimitsMemorySwappiness, state.LimitsMemorySwappiness),
		LimitsMemoryReservation: configure.DiffString(plan.LimitsMemoryReservation, state.LimitsMemoryReservation),
		LimitsCpus:              configure.DiffString(plan.LimitsCpus, state.LimitsCpus),
		LimitsCpuset:            configure.DiffString(plan.LimitsCpuset, state.LimitsCpuset),
		LimitsCPUShares:         configure.DiffInt64(plan.LimitsCPUShares, state.LimitsCPUShares),

		CustomLabels:           configure.DiffString(plan.CustomLabels, state.CustomLabels),
		CustomDockerRunOptions: configure.DiffString(plan.CustomDockerRunOptions, state.CustomDockerRunOptions),

		IsForceHTTPSEnabled:    configure.DiffBool(plan.IsForceHTTPSEnabled, state.IsForceHTTPSEnabled),
		ConnectToDockerNetwork: configure.DiffBool(plan.ConnectToDockerNetwork, state.ConnectToDockerNetwork),
	}

	if err := r.client.Application.Update(ctx, plan.Id.ValueString(), &updateDTO); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update docker image application, got error: %s", err))
		return
	}

	if !plan.EnvironmentVariables.Equal(state.EnvironmentVariables) {
		if diags := syncEnvs(ctx, r, plan.Id.ValueString(), plan.EnvironmentVariables, state.EnvironmentVariables); diags != nil {
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
		}
	}

	if plan.RedeployOnUpdate.ValueBool() && runtimeFieldsChanged(plan, state) {
		if _, err := r.client.Application.Restart(ctx, plan.Id.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Update succeeded but redeploy failed: %s", err))
			return
		}
	}

	app, err := r.client.Application.Get(ctx, plan.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read application after update, got error: %s", err))
		return
	}
	mapApplicationToModel(&plan, app)

	envs, err := r.client.Application.ListEnvs(ctx, plan.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list envs after update, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(mapEnvsToModel(ctx, &plan, *envs)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}
