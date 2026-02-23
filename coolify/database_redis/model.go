package database_redis

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/marconneves/coolify-sdk-go/database"
)

type DatabaseRedisModel struct {
	Id                      types.String `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	Description             types.String `tfsdk:"description"`
	ServerUUID              types.String `tfsdk:"server_uuid"`
	ProjectUUID             types.String `tfsdk:"project_uuid"`
	EnvironmentName         types.String `tfsdk:"environment_name"`
	DestinationUUID         types.String `tfsdk:"destination_uuid"`
	Image                   types.String `tfsdk:"image"`
	IsPublic                types.Bool   `tfsdk:"is_public"`
	PublicPort              types.Int64  `tfsdk:"public_port"`
	InstantDeploy           types.Bool   `tfsdk:"instant_deploy"`
	
	RedisPassword           types.String `tfsdk:"redis_password"`
	RedisConf               types.String `tfsdk:"redis_conf"`

	LimitsMemory            types.String `tfsdk:"limits_memory"`
	LimitsMemorySwap        types.String `tfsdk:"limits_memory_swap"`
	LimitsMemorySwappiness  types.Int64  `tfsdk:"limits_memory_swappiness"`
	LimitsMemoryReservation types.String `tfsdk:"limits_memory_reservation"`
	LimitsCPUs              types.String `tfsdk:"limits_cpus"`
	LimitsCPUSet            types.String `tfsdk:"limits_cpuset"`
	LimitsCPUShares         types.Int64  `tfsdk:"limits_cpu_shares"`

	Status                  types.String `tfsdk:"status"`
}

func mapRedisResourceModel(data *DatabaseRedisModel, db *database.Database) {
	data.Id = types.StringValue(db.UUID)
	data.Name = types.StringValue(db.Name)
	if db.Description != nil {
		data.Description = types.StringValue(*db.Description)
	} else {
		data.Description = types.StringNull()
	}
	data.Image = types.StringValue(db.Image)
	data.IsPublic = types.BoolValue(db.IsPublic)
	data.PublicPort = types.Int64Value(int64(db.PublicPort))
	
	if db.RedisPassword != nil {
		data.RedisPassword = types.StringValue(*db.RedisPassword)
	}
	if db.RedisConf != nil {
		data.RedisConf = types.StringValue(*db.RedisConf)
	}

	data.LimitsMemory = types.StringValue(db.LimitsMemory)
	data.LimitsMemorySwap = types.StringValue(db.LimitsMemorySwap)
	data.LimitsMemorySwappiness = types.Int64Value(int64(db.LimitsMemorySwappiness))
	data.LimitsMemoryReservation = types.StringValue(db.LimitsMemoryReservation)
	data.LimitsCPUs = types.StringValue(db.LimitsCpus)
	if db.LimitsCpuset != nil {
		data.LimitsCPUSet = types.StringValue(*db.LimitsCpuset)
	}
	data.LimitsCPUShares = types.Int64Value(int64(db.LimitsCPUShares))

	data.Status = types.StringValue(db.Status)
}
