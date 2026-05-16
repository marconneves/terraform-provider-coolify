package application_dockerimage

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/application"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// CreateApplication creates the docker image application and applies the
// environment variables map in bulk.
func (r *ApplicationDockerImageResource) CreateApplication(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApplicationDockerImageModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createDTO := application.CreateDockerImageDTO{
		ProjectUUID:             data.ProjectUUID.ValueString(),
		ServerUUID:              data.ServerUUID.ValueString(),
		EnvironmentName:         data.EnvironmentName.ValueString(),
		EnvironmentUUID:         data.EnvironmentUUID.ValueString(),
		DockerRegistryImageName: data.DockerRegistryImageName.ValueString(),
		PortsExposes:            data.PortsExposes.ValueString(),

		DockerRegistryImageTag: configure.ValueStringPointer(data.DockerRegistryImageTag),
		DestinationUUID:        configure.ValueStringPointer(data.DestinationUUID),
		Name:                   configure.ValueStringPointer(data.Name),
		Description:            configure.ValueStringPointer(data.Description),
		Domains:                configure.ValueStringPointer(data.Domains),
		PortsMappings:          configure.ValueStringPointer(data.PortsMappings),

		HealthCheckEnabled:      data.HealthCheckEnabled.ValueBoolPointer(),
		HealthCheckPath:         configure.ValueStringPointer(data.HealthCheckPath),
		HealthCheckPort:         configure.ValueStringPointer(data.HealthCheckPort),
		HealthCheckHost:         configure.ValueStringPointer(data.HealthCheckHost),
		HealthCheckMethod:       configure.ValueStringPointer(data.HealthCheckMethod),
		HealthCheckReturnCode:   configure.Int64ToUintPtr(data.HealthCheckReturnCode),
		HealthCheckScheme:       configure.ValueStringPointer(data.HealthCheckScheme),
		HealthCheckResponseText: configure.ValueStringPointer(data.HealthCheckResponseText),
		HealthCheckInterval:     configure.Int64ToUintPtr(data.HealthCheckInterval),
		HealthCheckTimeout:      configure.Int64ToUintPtr(data.HealthCheckTimeout),
		HealthCheckRetries:      configure.Int64ToUintPtr(data.HealthCheckRetries),
		HealthCheckStartPeriod:  configure.Int64ToUintPtr(data.HealthCheckStartPeriod),

		LimitsMemory:            configure.ValueStringPointer(data.LimitsMemory),
		LimitsMemorySwap:        configure.ValueStringPointer(data.LimitsMemorySwap),
		LimitsMemorySwappiness:  configure.Int64ToUintPtr(data.LimitsMemorySwappiness),
		LimitsMemoryReservation: configure.ValueStringPointer(data.LimitsMemoryReservation),
		LimitsCpus:              configure.ValueStringPointer(data.LimitsCpus),
		LimitsCpuset:            configure.ValueStringPointer(data.LimitsCpuset),
		LimitsCPUShares:         configure.Int64ToUintPtr(data.LimitsCPUShares),

		CustomLabels:           configure.ValueStringPointer(data.CustomLabels),
		CustomDockerRunOptions: configure.ValueStringPointer(data.CustomDockerRunOptions),

		InstantDeploy:          data.InstantDeploy.ValueBoolPointer(),
		IsForceHTTPSEnabled:    data.IsForceHTTPSEnabled.ValueBoolPointer(),
		ConnectToDockerNetwork: data.ConnectToDockerNetwork.ValueBoolPointer(),
	}

	uuid, err := r.client.Application.CreateDockerImage(ctx, &createDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create docker image application, got error: %s", err))
		return
	}

	if !data.EnvironmentVariables.IsNull() && !data.EnvironmentVariables.IsUnknown() && len(data.EnvironmentVariables.Elements()) > 0 {
		envs, diags := envDTOsFromMap(ctx, data.EnvironmentVariables)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		if _, err := r.client.Application.UpdateEnvsBulk(ctx, *uuid, envs); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Application %s created, but bulk env upsert failed: %s", *uuid, err))
			data.Id = stringValueFromPtr(uuid)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	app, err := r.client.Application.Get(ctx, *uuid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read application after creation, got error: %s", err))
		return
	}

	mapApplicationToModel(&data, app)

	envs, err := r.client.Application.ListEnvs(ctx, *uuid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list envs after creation, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(mapEnvsToModel(ctx, &data, *envs)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
