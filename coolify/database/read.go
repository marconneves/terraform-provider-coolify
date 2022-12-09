package database

import (
	"context"
	"strings"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func databaseReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	databaseId := d.Id()


	item, err := apiClient.GetDatabase(databaseId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return diag.Errorf("error finding Item with ID %v", databaseId)
		}
	}

	d.SetId(item.Database.Id)
	d.Set("name", item.Database.Name)

	engine := make(map[string]interface{})
	engine["image"] = item.Database.Type
	engine["version"] = item.Database.Version

	d.Set("engine", []interface{}{engine})

	settings := make(map[string]interface{})
	settings["destination_id"] = item.Database.DestinationDockerId
	settings["is_public"] = item.Database.Settings.IsPublic
	settings["append_only"] = item.Database.Settings.AppendOnly
	d.Set("settings", []interface{}{settings})

	
	status := make(map[string]interface{})
	if item.Database.Settings.IsPublic == true {
		status["port"] = item.Database.PublicPort
	} else {
		status["port"] = item.PrivatePort
	}
	
	d.Set("status", []interface{}{status})

	return nil
}