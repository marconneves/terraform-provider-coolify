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
		PublicPort:      configure.Int64ToUintPtr(data.PublicPort),
		InstantDeploy:   data.InstantDeploy.ValueBoolPointer(),
		DestinationUUID: configure.ValueStringPointer(data.DestinationUUID),

		MysqlRootPassword: configure.ValueStringPointer(data.MysqlRootPassword),
		MysqlPassword:     configure.ValueStringPointer(data.MysqlPassword),
		MysqlUser:         configure.ValueStringPointer(data.MysqlUser),
		MysqlDatabase:     configure.ValueStringPointer(data.MysqlDatabase),
		MysqlConf:         configure.Base64EncodeString(data.MysqlConf),

		LimitsMemory:            configure.ValueStringPointer(data.LimitsMemory),
		LimitsMemorySwap:        configure.ValueStringPointer(data.LimitsMemorySwap),
		LimitsMemorySwappiness:  configure.Int64ToUintPtr(data.LimitsMemorySwappiness),
		LimitsMemoryReservation: configure.ValueStringPointer(data.LimitsMemoryReservation),
		LimitsCPUs:              configure.ValueStringPointer(data.LimitsCPUs),
		LimitsCPUSet:            configure.ValueStringPointer(data.LimitsCPUSet),
		LimitsCPUShares:         configure.Int64ToUintPtr(data.LimitsCPUShares),
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
