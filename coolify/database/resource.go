package database

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/shared"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: databaseCreateItem,
		ReadContext: databaseReadItem,
		UpdateContext: databaseUpdateItem,
		DeleteContext: databaseDeleteItem,
		Exists: databaseExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Description:  "Name of the database.",
				Required:     true,
				ForceNew:     false,
				ValidateFunc: shared.ValidateName,
			},

			"engine": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"image": {
							Type:     schema.TypeString,
							Description: "Engine of db, options: MongoDB, MySQL, MariaDB, PostgreSQL, Redis, CouchDB or EdgeDB.",
							Required: true,
							ForceNew: true,
							ValidateFunc: ValidateEngine,
						},
						"version": {
							Type:          schema.TypeString,
							Required:      true,
							ForceNew:      true,
						},
					},
				},
			},

			"settings": {
				Type:     schema.TypeList,
				Required: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_id": {
							Type:          schema.TypeString,
							Required:      true,
							ForceNew:      true,
						},
						"is_public": {
							Type:     schema.TypeBool,
							Required: false,
							Default: false,
						},
						"append_only": {
							Type:     schema.TypeBool,
							Required: false,
							Default: false,
						},
					},
				},
			},
		},
	}
}


type Settings struct {
	public_port *int
}