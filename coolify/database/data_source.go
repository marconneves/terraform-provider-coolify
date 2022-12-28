package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceReadDatabase,
		Schema: map[string]*schema.Schema{
			"database_id": {
				Type:        schema.TypeString,
				Description: "Your database id.",
				Required:    true,
			},

			"name": {
				Type:        schema.TypeString,
				Description: "Name of the database.",
				Computed:    true,
			},

			"engine": {
				Type:        schema.TypeString,
				Description: "Engine of db, options: MongoDB, MySQL, MariaDB, PostgreSQL, Redis, CouchDB or EdgeDB with specific version.",
				Computed:    true,
			},

			"settings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_public": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"append_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"default_database": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"user": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"root_user": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"root_password": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
		UseJSONNumber: true,
	}
}

func dataSourceReadDatabase(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	databaseId := d.Get("database_id").(string)
	d.SetId(databaseId)

	return databaseReadItem(ctx, d, m)
}
