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
				Description:  "The name of the resource, also acts as it's unique ID",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: shared.ValidateName,
			},

			"engine": {
				Type:     schema.TypeString,
				Description: "Engine of db, options: MongoDB, MySQL, MariaDB, PostgreSQL, Redis, CouchDB or EdgeDB.",
				Required: true,
				ForceNew: true,
				ValidateFunc: ValidateEngine,
			},

			"engine_version": {
				Type:          schema.TypeString,
				Required:      true,
				ForceNew:      true,
			},

			"destination_id": {
				Type:          schema.TypeString,
				Required:      true,
				ForceNew:      true,
			},

			"is_public": {
				Type:        schema.TypeBool,
				Description: "If this database is public or not",
				Optional:    true,
				Default:    false,
			},
			
			"settings": {
				Type:        schema.TypeSet,
				Description: "Optional settings for the database",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}
