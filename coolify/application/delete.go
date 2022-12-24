package application

import (
	"context"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func applicationDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	itemId := d.Id()

	err := apiClient.StopApplication(itemId)
	if err != nil {
		return diag.FromErr(err)
	}

	err = apiClient.DeleteApplication(itemId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil

}
