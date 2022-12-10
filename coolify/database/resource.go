package database

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/shared"
)

type Database struct {
	Name string `json:"name"`
	Engine struct {
		Image string `json:"image"`
		Version string `json:"version"`
	} `json:"engine"`
	Settings struct {
		DestinationId string `json:"destination_id"`
		IsPublic bool `json:"is_public"`
		AppendOnly bool `json:"append_only"`
	} `json:"settings"`
	Status struct {
		PublicPort int `json:"public_port"`
	}
}

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
				Type:         schema.TypeString,
				Description:  "Engine of db, options: MongoDB, MySQL, MariaDB, PostgreSQL, Redis, CouchDB or EdgeDB with specific version.",
				Required:     true,
				ForceNew:     false,
				ValidateFunc: ValidateEngine,
			},

			"settings": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_id": {
							Type:          schema.TypeString,
							Required:      true,
							ForceNew:      true,
						},
						"is_public": {
							Type:     schema.TypeBool,
							Optional: true,
							Default: false,
						},
						"append_only": {
							Type:     schema.TypeBool,
							Optional: true,
							Default: false,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}