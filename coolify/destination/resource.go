package destination

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: destinationCreateItem,
	}
}