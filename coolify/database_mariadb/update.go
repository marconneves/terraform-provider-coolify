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
	plan := data
	var state DatabaseMariaDBModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateDTO := database.UpdateDatabaseDTO{
		Name:                    configure.DiffString(plan.Name, state.Name),
		Description:             configure.DiffString(plan.Description, state.Description),
		Image:                   configure.DiffString(plan.Image, state.Image),
		IsPublic:                configure.DiffBool(plan.IsPublic, state.IsPublic),
		PublicPort:              configure.DiffInt64(plan.PublicPort, state.PublicPort),

		MariadbRootPassword: configure.DiffString(plan.MariadbRootPassword, state.MariadbRootPassword),
		MariadbPassword:     configure.DiffString(plan.MariadbPassword, state.MariadbPassword),
		MariadbUser:         configure.DiffString(plan.MariadbUser, state.MariadbUser),
		MariadbDatabase:     configure.DiffString(plan.MariadbDatabase, state.MariadbDatabase),
		MariadbConf:         configure.DiffBase64String(plan.MariadbConf, state.MariadbConf),

		LimitsMemory:            configure.DiffString(plan.LimitsMemory, state.LimitsMemory),
		LimitsMemorySwap:        configure.DiffString(plan.LimitsMemorySwap, state.LimitsMemorySwap),
		LimitsMemorySwappiness:  configure.DiffInt64(plan.LimitsMemorySwappiness, state.LimitsMemorySwappiness),
		LimitsMemoryReservation: configure.DiffString(plan.LimitsMemoryReservation, state.LimitsMemoryReservation),
		LimitsCpus:              configure.DiffString(plan.LimitsCPUs, state.LimitsCPUs),
		LimitsCpuset:            configure.DiffString(plan.LimitsCPUSet, state.LimitsCPUSet),
		LimitsCPUShares:         configure.DiffInt64(plan.LimitsCPUShares, state.LimitsCPUShares),
	}

	err := r.client.Database.Update(ctx, plan.Id.ValueString(), &updateDTO)
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
