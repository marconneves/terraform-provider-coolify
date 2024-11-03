package destination

import (
	"context"
	"strings"

	sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func destinationReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*sdk.Client)
	destinationId := d.Id()

	destination, err := apiClient.GetDestination(destinationId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return diag.Errorf("error finding Item with ID %v", destinationId)
		}
	}

	d.SetId(destination.Id)

	d.Set("name", destination.Name)
	d.Set("engine", destination.Engine)
	d.Set("network", destination.Network)

	return nil
}
