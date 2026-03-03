package database_mariadb

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marconneves/coolify-sdk-go/database"
	configure "github.com/marconneves/terraform-provider-coolify/shared"
)

// CreateMariaDB creates a new MariaDB database resource.
func (r *MariaDBResource) CreateMariaDB(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DatabaseMariaDBModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createDTO := database.CreateDatabaseMariaDBDTO{
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

		MariadbRootPassword: configure.ValueStringPointer(data.MariadbRootPassword),
		MariadbPassword:     configure.ValueStringPointer(data.MariadbPassword),
		MariadbUser:         configure.ValueStringPointer(data.MariadbUser),
		MariadbDatabase:     configure.ValueStringPointer(data.MariadbDatabase),
		MariadbConf:         configure.ValueStringPointer(data.MariadbConf),

		LimitsMemory:            configure.ValueStringPointer(data.LimitsMemory),
		LimitsMemorySwap:        configure.ValueStringPointer(data.LimitsMemorySwap),
		LimitsMemorySwappiness:  configure.Int64ToUintPtr(data.LimitsMemorySwappiness),
		LimitsMemoryReservation: configure.ValueStringPointer(data.LimitsMemoryReservation),
		LimitsCPUs:              configure.ValueStringPointer(data.LimitsCPUs),
		LimitsCPUSet:            configure.ValueStringPointer(data.LimitsCPUSet),
		LimitsCPUShares:         configure.Int64ToUintPtr(data.LimitsCPUShares),
	}

	uuid, err := r.client.Database.CreateMariaDB(ctx, &createDTO)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create MariaDB database, got error: %s", err))
		return
	}

	db, err := r.client.Database.Get(ctx, *uuid)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read MariaDB database after creation, got error: %s", err))
		return
	}

	mapMariaDBResourceModel(&data, db)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
