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
		Name:                    data.Name.ValueStringPointer(),
		Description:             data.Description.ValueStringPointer(),
		Image:                   data.Image.ValueStringPointer(),
		IsPublic:                data.IsPublic.ValueBoolPointer(),
		PublicPort:              configure.Int64ToUintPtr(data.PublicPort.ValueInt64Pointer()),

		MysqlRootPassword: data.MysqlRootPassword.ValueStringPointer(),
		MysqlPassword:     data.MysqlPassword.ValueStringPointer(),
		MysqlUser:         data.MysqlUser.ValueStringPointer(),
		MysqlDatabase:     data.MysqlDatabase.ValueStringPointer(),
		MysqlConf:         data.MysqlConf.ValueStringPointer(),

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
