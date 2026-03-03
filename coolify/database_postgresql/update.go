package database_postgresql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/database"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// UpdatePostgres updates a PostgreSQL database resource.
func (r *PostgresResource) UpdatePostgres(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DatabasePostgresModel

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

		PostgresUser:           configure.ValueStringPointer(data.PostgresUser),
		PostgresPassword:       configure.ValueStringPointer(data.PostgresPassword),
		PostgresDB:             configure.ValueStringPointer(data.PostgresDB),
		PostgresInitdbArgs:     configure.ValueStringPointer(data.PostgresInitdbArgs),
		PostgresHostAuthMethod: configure.ValueStringPointer(data.PostgresHostAuthMethod),
		PostgresConf:           configure.ValueStringPointer(data.PostgresConf),

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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update PostgreSQL database, got error: %s", err))
		return
	}

	db, err := r.client.Database.Get(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read PostgreSQL database after update, got error: %s", err))
		return
	}

	mapPostgresResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
