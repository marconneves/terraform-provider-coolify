package application

import (
	"context"

	sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func applicationDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*sdk.Client)
	itemId := d.Id()

	err := apiClient.StopApplication(itemId)
	if err != nil {
		return diag.FromErr(err)
	}

	err = apiClient.DeleteApplication(itemId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil

}
