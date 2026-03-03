package database_redis

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/database"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// CreateRedis creates a new Redis database resource.
func (r *RedisResource) CreateRedis(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DatabaseRedisModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createDTO := database.CreateDatabaseRedisDTO{
		ServerUUID:      data.ServerUUID.ValueString(),
		ProjectUUID:     data.ProjectUUID.ValueString(),
		EnvironmentName: data.EnvironmentName.ValueString(),
		Name:            data.Name.ValueStringPointer(),
		Description:     data.Description.ValueStringPointer(),
		Image:           data.Image.ValueStringPointer(),
		IsPublic:        data.IsPublic.ValueBoolPointer(),
		PublicPort:      configure.Int64ToUintPtr(data.PublicPort),
		InstantDeploy:   data.InstantDeploy.ValueBoolPointer(),
		DestinationUUID: configure.ValueStringPointer(data.DestinationUUID),

		RedisPassword: configure.ValueStringPointer(data.RedisPassword),
		RedisConf:     configure.ValueStringPointer(data.RedisConf),

		LimitsMemory:            configure.ValueStringPointer(data.LimitsMemory),
		LimitsMemorySwap:        configure.ValueStringPointer(data.LimitsMemorySwap),
		LimitsMemorySwappiness:  configure.Int64ToUintPtr(data.LimitsMemorySwappiness),
		LimitsMemoryReservation: configure.ValueStringPointer(data.LimitsMemoryReservation),
		LimitsCPUs:              configure.ValueStringPointer(data.LimitsCPUs),
		LimitsCPUSet:            configure.ValueStringPointer(data.LimitsCPUSet),
		LimitsCPUShares:         configure.Int64ToUintPtr(data.LimitsCPUShares),
	}

	uuid, err := r.client.Database.CreateRedis(ctx, &createDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Redis database, got error: %s", err))
		return
	}

	db, err := r.client.Database.Get(ctx, *uuid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Redis database after creation, got error: %s", err))
		return
	}

	mapRedisResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
