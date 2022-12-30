package database

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/shared"
)

type Status struct {
	Port string `json:"port"`
}

type Database struct {
	Name   string
	Engine struct {
		Image   string
		Version string
	}
	Settings struct {
		DestinationId   string
		IsPublic        bool
		AppendOnly      bool
		DefaultDatabase string
		User            string
		Password        string
		RootUser        string
		RootPassword    string
	}
	Status Status
}

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: databaseCreateItem,
		ReadContext:   databaseRead,
		UpdateContext: databaseUpdateItem,
		DeleteContext: databaseDeleteItem,
		Exists:        databaseExistsItem,
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
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"is_public": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
							Default:  true,
						},
						"append_only": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"default_database": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Default:  "",
						},
						"user": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Default:  "",
						},
						"password": {
							Type:      schema.TypeString,
							Optional:  true,
							ForceNew:  true,
							Sensitive: true,
							Default:   "",
						},
						"root_user": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Default:  "",
						},
						"root_password": {
							Type:      schema.TypeString,
							Optional:  true,
							ForceNew:  true,
							Sensitive: true,
							Default:   "",
						},
					},
				},
			},

			"uri": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"status": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}
