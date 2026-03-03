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
		PublicPort:      configure.Int64ToUintPtr(data.PublicPort),
		InstantDeploy:   data.InstantDeploy.ValueBoolPointer(),
		DestinationUUID: configure.ValueStringPointer(data.DestinationUUID),

		PostgresUser:           configure.ValueStringPointer(data.PostgresUser),
		PostgresPassword:       configure.ValueStringPointer(data.PostgresPassword),
		PostgresDB:             configure.ValueStringPointer(data.PostgresDB),
		PostgresInitdbArgs:     configure.ValueStringPointer(data.PostgresInitdbArgs),
		PostgresHostAuthMethod: configure.ValueStringPointer(data.PostgresHostAuthMethod),
		PostgresConf:           configure.Base64EncodeString(data.PostgresConf),

		LimitsMemory:            configure.ValueStringPointer(data.LimitsMemory),
		LimitsMemorySwap:        configure.ValueStringPointer(data.LimitsMemorySwap),
		LimitsMemorySwappiness:  configure.Int64ToUintPtr(data.LimitsMemorySwappiness),
		LimitsMemoryReservation: configure.ValueStringPointer(data.LimitsMemoryReservation),
		LimitsCPUs:              configure.ValueStringPointer(data.LimitsCPUs),
		LimitsCPUSet:            configure.ValueStringPointer(data.LimitsCPUSet),
		LimitsCPUShares:         configure.Int64ToUintPtr(data.LimitsCPUShares),
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
