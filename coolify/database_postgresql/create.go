package database_postgresql

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/database"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// CreatePostgres creates a new PostgreSQL database resource.
func (r *PostgresResource) CreatePostgres(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DatabasePostgresModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createDTO := database.CreateDatabasePostgresDTO{
		ServerUUID:      data.ServerUUID.ValueString(),
		ProjectUUID:     data.ProjectUUID.ValueString(),
		Environment:     data.EnvironmentName.ValueString(),
		Name:            data.Name.ValueStringPointer(),
		Description:     data.Description.ValueStringPointer(),
		Image:           data.Image.ValueStringPointer(),
		IsPublic:        data.IsPublic.ValueBoolPointer(),
		PublicPort:      configure.Int64ToUintPtr(data.PublicPort.ValueInt64Pointer()),
		InstantDeploy:   data.InstantDeploy.ValueBoolPointer(),
		DestinationUUID: data.DestinationUUID.ValueStringPointer(),

		PostgresUser:           data.PostgresUser.ValueStringPointer(),
		PostgresPassword:       data.PostgresPassword.ValueStringPointer(),
		PostgresDB:             data.PostgresDB.ValueStringPointer(),
		PostgresInitdbArgs:     data.PostgresInitdbArgs.ValueStringPointer(),
		PostgresHostAuthMethod: data.PostgresHostAuthMethod.ValueStringPointer(),
		PostgresConf:           data.PostgresConf.ValueStringPointer(),

		LimitsMemory:            data.LimitsMemory.ValueStringPointer(),
		LimitsMemorySwap:        data.LimitsMemorySwap.ValueStringPointer(),
		LimitsMemorySwappiness:  configure.Int64ToUintPtr(data.LimitsMemorySwappiness.ValueInt64Pointer()),
		LimitsMemoryReservation: data.LimitsMemoryReservation.ValueStringPointer(),
		LimitsCPUs:              data.LimitsCPUs.ValueStringPointer(),
		LimitsCPUSet:            data.LimitsCPUSet.ValueStringPointer(),
		LimitsCPUShares:         configure.Int64ToUintPtr(data.LimitsCPUShares.ValueInt64Pointer()),
	}

	uuid, err := r.client.Database.CreatePostgreSQL(ctx, &createDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create PostgreSQL database, got error: %s", err))
		return
	}

	db, err := r.client.Database.Get(ctx, *uuid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read PostgreSQL database after creation, got error: %s", err))
		return
	}

	mapPostgresResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
