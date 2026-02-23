package database_mysql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/database"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// CreateMySQL creates a new MySQL database resource.
func (r *MySQLResource) CreateMySQL(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DatabaseMySQLModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createDTO := database.CreateDatabaseMySQLDTO{
		ServerUUID:      data.ServerUUID.ValueString(),
		ProjectUUID:     data.ProjectUUID.ValueString(),
		EnvironmentName: data.EnvironmentName.ValueString(),
		EnvironmentUUID: data.EnvironmentUUID.ValueStringPointer(),
		Name:            data.Name.ValueStringPointer(),
		Description:     data.Description.ValueStringPointer(),
		Image:           data.Image.ValueStringPointer(),
		IsPublic:        data.IsPublic.ValueBoolPointer(),
		PublicPort:      configure.Int64ToUintPtr(data.PublicPort.ValueInt64Pointer()),
		InstantDeploy:   data.InstantDeploy.ValueBoolPointer(),
		DestinationUUID: data.DestinationUUID.ValueStringPointer(),

		MysqlRootPassword: data.MysqlRootPassword.ValueStringPointer(),
		MysqlPassword:     data.MysqlPassword.ValueStringPointer(),
		MysqlUser:         data.MysqlUser.ValueStringPointer(),
		MysqlDatabase:     data.MysqlDatabase.ValueStringPointer(),
		MysqlConf:         data.MysqlConf.ValueStringPointer(),

		LimitsMemory:            data.LimitsMemory.ValueStringPointer(),
		LimitsMemorySwap:        data.LimitsMemorySwap.ValueStringPointer(),
		LimitsMemorySwappiness:  configure.Int64ToUintPtr(data.LimitsMemorySwappiness.ValueInt64Pointer()),
		LimitsMemoryReservation: data.LimitsMemoryReservation.ValueStringPointer(),
		LimitsCPUs:              data.LimitsCPUs.ValueStringPointer(),
		LimitsCPUSet:            data.LimitsCPUSet.ValueStringPointer(),
		LimitsCPUShares:         configure.Int64ToUintPtr(data.LimitsCPUShares.ValueInt64Pointer()),
	}

	uuid, err := r.client.Database.CreateMySQL(ctx, &createDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create MySQL database, got error: %s", err))
		return
	}

	db, err := r.client.Database.Get(ctx, *uuid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read MySQL database after creation, got error: %s", err))
		return
	}

	mapMySQLResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
