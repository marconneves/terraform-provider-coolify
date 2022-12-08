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

	databaseToUpdate := &client.UpdateDatabaseDTO{
		Name: d.Get("name").(string),
	}
	err := apiClient.UpdateDatabase(databaseId, databaseToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	settingsToUpdate := &client.UpdateSettingsDatabaseDTO{
		IsPublic: d.Get("is_public").(bool),
	}
	settingsResponse, err := apiClient.UpdateSettings(databaseId, settingsToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	if settingsResponse.PublicPort != nil {
		publicPort := strconv.Itoa(*settingsResponse.PublicPort)
		
		d.Set("settings.public_port", publicPort)
					
		tflog.Trace(ctx, "Database %v started on port: %" + databaseId + publicPort)
	}
	
	return nil
}