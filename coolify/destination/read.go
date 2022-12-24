package destination

import (
	"context"
	"strings"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func destinationReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	destinationId := d.Id()

	item, err := apiClient.GetDestination(destinationId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return diag.Errorf("error finding Item with ID %v", destinationId)
		}
	}

	d.SetId(item.Destination.Id)
	d.Set("name", item.Destination.Name)
	d.Set("engine", item.Destination.Engine)

	status := make(map[string]string)
	status["network"] = item.Destination.Network
	d.Set("status", status)

	return nil
}
