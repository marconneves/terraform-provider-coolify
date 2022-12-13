package destination

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func destinationExistItem(d *schema.ResourceData, m interface{}) (bool, error) {
	return true, nil
}