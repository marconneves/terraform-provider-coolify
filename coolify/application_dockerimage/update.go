package application_dockerimage

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/application"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

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
