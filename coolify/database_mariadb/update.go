package database_mariadb

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/database"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// UpdateMariaDB updates a MariaDB database resource.
func (r *MariaDBResource) UpdateMariaDB(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DatabaseMariaDBModel

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

		MariadbRootPassword: data.MariadbRootPassword.ValueStringPointer(),
		MariadbPassword:     data.MariadbPassword.ValueStringPointer(),
		MariadbUser:         data.MariadbUser.ValueStringPointer(),
		MariadbDatabase:     data.MariadbDatabase.ValueStringPointer(),
		MariadbConf:         data.MariadbConf.ValueStringPointer(),

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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update MariaDB database, got error: %s", err))
		return
	}

	db, err := r.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read MariaDB database after update, got error: %s", err))
		return
	}

	mapMariaDBResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
