package database

import (
	"context"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func databaseDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	itemId := d.Id()

	err := apiClient.StopDatabase(itemId)
	if err != nil {
		return diag.FromErr(err)
	}

	err = apiClient.DeleteDatabase(itemId, false)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
