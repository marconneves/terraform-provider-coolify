package destination

import (
	"strings"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func destinationExistItem(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)
	destinationId := d.Id()

	_, err := apiClient.GetDestination(destinationId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
