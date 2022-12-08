package database

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/api/client"
)



func databaseCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	id, err := apiClient.NewDatabase()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*id)

	err = apiClient.SetEngineDatabase(*id, d.Get("engine").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	err = apiClient.SetDestinationDatabase(*id, d.Get("destination_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	databaseToUpdate := &client.UpdateDatabaseDTO{
		Name: d.Get("name").(string),
		Version: d.Get("engine_version").(string),
		DefaultDatabase: "fist-db",
		DbUser: "user",
		DbUserPassword: "password",
	}

	err = apiClient.UpdateDatabase(*id, databaseToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}


	tflog.Trace(ctx, "Starting database...")
	err = apiClient.StartDatabase(*id)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Trace(ctx, "Data base started")

	if d.Get("is_public") != nil {
		settingsToUpdate := &client.UpdateSettingsDatabaseDTO{
			IsPublic: d.Get("is_public").(bool),
		}
		settingsResponse, err := apiClient.UpdateSettings(*id, settingsToUpdate)
		if err != nil {
			return diag.FromErr(err)
		}
	
		if settingsResponse.PublicPort != nil {
			publicPort := strconv.Itoa(*settingsResponse.PublicPort)
			
			d.Set("settings.public_port", publicPort)
						
			tflog.Trace(ctx, "Database %v started on port: %" + *id + publicPort)
		}
	}
	
	return nil
}