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
	status := make(map[string]interface{})
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

		db.Settings.AppendOnly = i["append_only"].(bool)
		db.Settings.DestinationId = i["destination_id"].(string)
		db.Settings.IsPublic = i["is_public"].(bool)
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

	settingsToUpdate := &client.UpdateSettingsDatabaseDTO{
		IsPublic: db.Settings.IsPublic,
	}
	settingsResponse, err := apiClient.UpdateSettings(*id, settingsToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	if settingsResponse.PublicPort != nil {
		status["port"] = settingsResponse.PublicPort
					
		tflog.Info(ctx, "Database %v started on port: %" + *id + strconv.Itoa(*settingsResponse.PublicPort))
	}
	
	d.Set("status", []interface{}{status})
	
	return nil
}