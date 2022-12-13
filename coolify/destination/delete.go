package destination

import (
	"context"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func destinationDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	destinationId := d.Id()

	err := apiClient.DeleteDestination(destinationId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}