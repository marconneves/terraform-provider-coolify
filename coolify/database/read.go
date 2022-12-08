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
	d.Set("engine", item.Database.Type)
	d.Set("engine_version", item.Database.Version)
	d.Set("destination_id", item.Database.DestinationDockerId)
	d.Set("is_public", item.Database.Settings.IsPublic)
	d.Set("settings", nil)
	return nil
}
