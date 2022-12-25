package destination

import (
	"context"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceReadItem,
		Schema: map[string]*schema.Schema{
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func dataSourceReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	network := d.Get("network").(string)

	destinations, err := apiClient.GetDestinations()
	if err != nil {
		return diag.Errorf("error on get network with name: %v", network)
	}

	fond := false

	for _, destination := range *destinations {
		if destination.Network == network {
			d.SetId(destination.Id)
			d.Set("name", destination.Name)
			d.Set("engine", destination.Engine)
			d.Set("network", destination.Network)
			fond = true
		}
	}

	if !fond {
		return diag.Errorf("error on get network with name: %v", network)
	}

	return nil
}
