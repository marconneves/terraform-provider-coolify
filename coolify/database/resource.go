package database

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/api/client"
	"terraform-provider-coolify/shared"
)

func Resource() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
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
		Create: databaseCreateItem,
		Read:   resourceReadItem,
		Update: resourceUpdateItem,
		Delete: resourceDeleteItem,
		Exists: resourceExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item, err := apiClient.GetItem(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Item with ID %s", itemId)
		}
	}

	d.SetId(item.Name)
	d.Set("name", item.Name)
	d.Set("description", item.Description)
	d.Set("tags", item.Tags)
	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	tfTags := d.Get("tags").(*schema.Set).List()
	tags := make([]string, len(tfTags))
	for i, tfTag := range tfTags {
		tags[i] = tfTag.(string)
	}

	item := client.Item{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Tags:        tags,
	}

	err := apiClient.UpdateItem(&item)
	if err != nil {
		return err
	}
	return nil
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteItem(itemId)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	_, err := apiClient.GetItem(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}