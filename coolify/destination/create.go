package destination

import (
	"context"
	"fmt"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func destinationCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// 1. Generate network id random
	// // 2. Generate new and get id
	// 3. Check if this network exist
	// 4. Update and set engine and network id

	apiClient := m.(*client.Client)

	destination := &client.CreateDestinationDTO{
		Name: "network",
		Network: "network",
	}
	destinationId, err := apiClient.NewDestination(destination)
	if err != nil {
		return diag.FromErr(err)
	}

	fmt.Printf("destination id: %v", *destinationId);


	return nil
}