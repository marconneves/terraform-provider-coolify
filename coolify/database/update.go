package database

import (
	"context"
	"strconv"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func databaseUpdateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	databaseId := d.Id()
	db := &Database{}


	err := apiClient.UpdateNameDatabase(databaseId, d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	settings := d.Get("settings").([]interface{})
	for _, setting := range settings {
		i := setting.(map[string]interface{})

		db.Settings.DestinationId = i["destination_id"].(string)
		db.Settings.IsPublic = i["is_public"].(bool)
		db.Settings.AppendOnly = i["append_only"].(bool)
	}

	settingsToUpdate := &client.UpdateSettingsDatabaseDTO{
		IsPublic: db.Settings.IsPublic,
	}
	settingsResponse, err := apiClient.UpdateSettings(databaseId, settingsToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	if settingsResponse.PublicPort != nil {
		settingsToSet := &Settings{
			public_port: settingsResponse.PublicPort,
		}
		// d.Set("settings", settingsToSet)
		// TODO: Set port after update
					
		tflog.Trace(ctx, "Database %v started on port: %" + databaseId + strconv.Itoa(*settingsToSet.public_port))
	}
	
	return nil
}