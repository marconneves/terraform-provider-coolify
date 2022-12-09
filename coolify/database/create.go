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

	db := &Database{}

	db.Name = d.Get("name").(string)


	engines := d.Get("engine").([]interface{})
	for _, engine := range engines {
		i := engine.(map[string]interface{})

		db.Engine.Image = i["image"].(string)
		db.Engine.Version = i["version"].(string)
	}

	settings := d.Get("settings").([]interface{})
	for _, setting := range settings {
		i := setting.(map[string]interface{})

		db.Settings.DestinationId = i["destination_id"].(string)
		db.Settings.IsPublic = i["is_public"].(bool)
		db.Settings.AppendOnly = i["append_only"].(bool)
	}

	apiClient := m.(*client.Client)

	id, err := apiClient.NewDatabase()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*id)

	err = apiClient.SetEngineDatabase(*id, db.Engine.Image)
	if err != nil {
		return diag.FromErr(err)
	}

	err = apiClient.SetDestinationDatabase(*id, db.Settings.DestinationId)
	if err != nil {
		return diag.FromErr(err)
	}

	databaseToUpdate := &client.UpdateDatabaseDTO{
		Name:  db.Name,
		Version: db.Engine.Version,
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
			IsPublic: db.Settings.IsPublic,
		}
		settingsResponse, err := apiClient.UpdateSettings(*id, settingsToUpdate)
		if err != nil {
			return diag.FromErr(err)
		}
	
		if settingsResponse.PublicPort != nil {
			settingsToSet := &Settings{
				public_port: settingsResponse.PublicPort,
			}
			// d.Set("settings", settingsToSet)
			// TODO: Set port after create
						
			tflog.Trace(ctx, "Database %v started on port: %" + *id + strconv.Itoa(*settingsToSet.public_port))
		}
	}
	
	return nil
}