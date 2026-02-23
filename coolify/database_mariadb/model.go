package database_mariadb

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/marconneves/coolify-sdk-go/database"
)

// DatabaseMariaDBModel represents the data model for a MariaDB database.
type DatabaseMariaDBModel struct {
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

	MariadbConf             types.String `tfsdk:"mariadb_conf"`
	MariadbRootPassword     types.String `tfsdk:"mariadb_root_password"`
	MariadbUser             types.String `tfsdk:"mariadb_user"`
	MariadbPassword         types.String `tfsdk:"mariadb_password"`
	MariadbDatabase         types.String `tfsdk:"mariadb_database"`

	LimitsMemory            types.String `tfsdk:"limits_memory"`
	LimitsMemorySwap        types.String `tfsdk:"limits_memory_swap"`
	LimitsMemorySwappiness  types.Int64  `tfsdk:"limits_memory_swappiness"`
	LimitsMemoryReservation types.String `tfsdk:"limits_memory_reservation"`
	LimitsCPUs              types.String `tfsdk:"limits_cpus"`
	LimitsCPUSet            types.String `tfsdk:"limits_cpuset"`
	LimitsCPUShares         types.Int64  `tfsdk:"limits_cpu_shares"`

	Status                  types.String `tfsdk:"status"`
}

func mapMariaDBResourceModel(data *DatabaseMariaDBModel, db *database.Database) {
	data.Id = types.StringValue(db.UUID)
	data.Name = types.StringValue(db.Name)
	if db.Description != nil {
		data.Description = types.StringValue(*db.Description)
	} else {
		data.Description = types.StringNull()
	}
	data.Image = types.StringValue(db.Image)
	data.Status = types.StringValue(db.Status)
	data.IsPublic = types.BoolValue(db.IsPublic)
	data.PublicPort = types.Int64Value(int64(db.PublicPort))

	if db.MariadbConf != nil {
		data.MariadbConf = types.StringValue(*db.MariadbConf)
	}
	if db.MariadbRootPassword != nil {
		data.MariadbRootPassword = types.StringValue(*db.MariadbRootPassword)
	}
	if db.MariadbUser != nil {
		data.MariadbUser = types.StringValue(*db.MariadbUser)
	}
	if db.MariadbPassword != nil {
		data.MariadbPassword = types.StringValue(*db.MariadbPassword)
	}
	if db.MariadbDatabase != nil {
		data.MariadbDatabase = types.StringValue(*db.MariadbDatabase)
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
}
