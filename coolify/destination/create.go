package destination

import (
	"context"
	"terraform-provider-coolify/api/client"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func destinationCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	status := make(map[string]string)
	
	networkId := d.Get("network").(string)
	if networkId == "" {
		networkId = uuid.New().String()
	}

	apiClient := m.(*client.Client)

	networkAlreadyExist := apiClient.CheckIfNetworkNameExist(networkId)
	if networkAlreadyExist == true {
		return diag.Errorf("This network already exist. Got %v", networkId)
	}

	destination := &client.CreateDestinationDTO{
		Name: d.Get("name").(string),
		Network: networkId,
		Engine: d.Get("engine").(string),
		RemoteEngine: false,
		IsCoolifyProxyUsed: true,
	}
	destinationId, err := apiClient.NewDestination(destination)
	if err != nil {
		return diag.FromErr(err)
	}

	status["network"] = networkId
	d.Set("status", status)

	d.SetId(*destinationId)
	
	return nil
}