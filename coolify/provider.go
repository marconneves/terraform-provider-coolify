package coolify

import (
	// "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/coolify/application"
	"terraform-provider-coolify/coolify/database"
	"terraform-provider-coolify/coolify/destination"

	sdk "github.com/marconneves/coolify-sdk-go"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_ADDRESS", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_TOKEN", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"coolify_destination": destination.Resource(),
			"coolify_database":    database.Resource(),
			"coolify_application": application.Resource(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"coolify_destination": destination.DataSource(),
			"coolify_database":    database.DataSource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := d.Get("address").(string)
	token := d.Get("token").(string)

	return sdk.NewClient(address, token), nil
}
