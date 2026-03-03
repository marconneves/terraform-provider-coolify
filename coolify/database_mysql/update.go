package database_mysql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/database"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// UpdateMySQL updates a MySQL database resource.
func (r *MySQLResource) UpdateMySQL(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DatabaseMySQLModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateDTO := database.UpdateDatabaseDTO{
		Name:                    configure.ValueStringPointer(data.Name),
		Description:             configure.ValueStringPointer(data.Description),
		Image:                   configure.ValueStringPointer(data.Image),
		IsPublic:                data.IsPublic.ValueBoolPointer(),
		PublicPort:              configure.Int64ToUintPtr(data.PublicPort),

		MysqlRootPassword: configure.ValueStringPointer(data.MysqlRootPassword),
		MysqlPassword:     configure.ValueStringPointer(data.MysqlPassword),
		MysqlUser:         configure.ValueStringPointer(data.MysqlUser),
		MysqlDatabase:     configure.ValueStringPointer(data.MysqlDatabase),
		MysqlConf:         configure.ValueStringPointer(data.MysqlConf),

		LimitsMemory:            configure.ValueStringPointer(data.LimitsMemory),
		LimitsMemorySwap:        configure.ValueStringPointer(data.LimitsMemorySwap),
		LimitsMemorySwappiness:  configure.Int64ToUintPtr(data.LimitsMemorySwappiness),
		LimitsMemoryReservation: configure.ValueStringPointer(data.LimitsMemoryReservation),
		LimitsCpus:              configure.ValueStringPointer(data.LimitsCPUs),
		LimitsCpuset:            configure.ValueStringPointer(data.LimitsCPUSet),
		LimitsCPUShares:         configure.Int64ToUintPtr(data.LimitsCPUShares),
	}

	err := r.client.Database.Update(ctx, data.Id.ValueString(), &updateDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update MySQL database, got error: %s", err))
		return
	}

	db, err := r.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read MySQL database after update, got error: %s", err))
		return
	}

	mapMySQLResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
