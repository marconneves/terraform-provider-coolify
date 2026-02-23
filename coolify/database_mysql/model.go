package database_mysql

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/marconneves/coolify-sdk-go/database"
)

type DatabaseMySQLModel struct {
	Id                      types.String `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	Description             types.String `tfsdk:"description"`
	ServerUUID              types.String `tfsdk:"server_uuid"`
	ProjectUUID             types.String `tfsdk:"project_uuid"`
	EnvironmentName         types.String `tfsdk:"environment_name"`
	EnvironmentUUID         types.String `tfsdk:"environment_uuid"`
	DestinationUUID         types.String `tfsdk:"destination_uuid"`
	Image                   types.String `tfsdk:"image"`
	IsPublic                types.Bool   `tfsdk:"is_public"`
	PublicPort              types.Int64  `tfsdk:"public_port"`
	InstantDeploy           types.Bool   `tfsdk:"instant_deploy"`
	
	MysqlRootPassword       types.String `tfsdk:"mysql_root_password"`
	MysqlPassword           types.String `tfsdk:"mysql_password"`
	MysqlUser               types.String `tfsdk:"mysql_user"`
	MysqlDatabase           types.String `tfsdk:"mysql_database"`
	MysqlConf               types.String `tfsdk:"mysql_conf"`

	LimitsMemory            types.String `tfsdk:"limits_memory"`
	LimitsMemorySwap        types.String `tfsdk:"limits_memory_swap"`
	LimitsMemorySwappiness  types.Int64  `tfsdk:"limits_memory_swappiness"`
	LimitsMemoryReservation types.String `tfsdk:"limits_memory_reservation"`
	LimitsCPUs              types.String `tfsdk:"limits_cpus"`
	LimitsCPUSet            types.String `tfsdk:"limits_cpuset"`
	LimitsCPUShares         types.Int64  `tfsdk:"limits_cpu_shares"`

	Status                  types.String `tfsdk:"status"`
}

func mapMySQLResourceModel(data *DatabaseMySQLModel, db *database.Database) {
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
	
	if db.MysqlRootPassword != nil {
		data.MysqlRootPassword = types.StringValue(*db.MysqlRootPassword)
	}
	if db.MysqlPassword != nil {
		data.MysqlPassword = types.StringValue(*db.MysqlPassword)
	}
	if db.MysqlUser != nil {
		data.MysqlUser = types.StringValue(*db.MysqlUser)
	}
	if db.MysqlDatabase != nil {
		data.MysqlDatabase = types.StringValue(*db.MysqlDatabase)
	}
	if db.MysqlConf != nil {
		data.MysqlConf = types.StringValue(*db.MysqlConf)
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
