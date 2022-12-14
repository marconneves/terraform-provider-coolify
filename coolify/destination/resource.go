package destination

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: destinationCreateItem,
		ReadContext: destinationReadItem,
		UpdateContext: destinationUpdateItem,
		DeleteContext: destinationDeleteItem,
		Exists: destinationExistItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"engine": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default: "/var/run/docker.sock",
			},

			"status": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}