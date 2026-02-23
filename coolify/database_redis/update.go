package database_redis

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/database"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// UpdateRedis updates a Redis database resource.
func (r *RedisResource) UpdateRedis(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DatabaseRedisModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateDTO := database.UpdateDatabaseDTO{
		Name:                    data.Name.ValueStringPointer(),
		Description:             data.Description.ValueStringPointer(),
		Image:                   data.Image.ValueStringPointer(),
		IsPublic:                data.IsPublic.ValueBoolPointer(),
		PublicPort:              configure.Int64ToUintPtr(data.PublicPort.ValueInt64Pointer()),

		RedisPassword: data.RedisPassword.ValueStringPointer(),
		RedisConf:     data.RedisConf.ValueStringPointer(),

		LimitsMemory:            data.LimitsMemory.ValueStringPointer(),
		LimitsMemorySwap:        data.LimitsMemorySwap.ValueStringPointer(),
		LimitsMemorySwappiness:  configure.Int64ToUintPtr(data.LimitsMemorySwappiness.ValueInt64Pointer()),
		LimitsMemoryReservation: data.LimitsMemoryReservation.ValueStringPointer(),
		LimitsCpus:              data.LimitsCPUs.ValueStringPointer(),
		LimitsCpuset:            data.LimitsCPUSet.ValueStringPointer(),
		LimitsCPUShares:         configure.Int64ToUintPtr(data.LimitsCPUShares.ValueInt64Pointer()),
	}

	err := r.client.Database.Update(ctx, data.Id.ValueString(), &updateDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Redis database, got error: %s", err))
		return
	}

	db, err := r.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Redis database after update, got error: %s", err))
		return
	}

	mapRedisResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
