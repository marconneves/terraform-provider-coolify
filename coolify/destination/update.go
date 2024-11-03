package destination

import (
	"context"

	sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func destinationUpdateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*sdk.Client)
	databaseId := d.Id()

	err := apiClient.UpdateNameDestination(databaseId, d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	return destinationReadItem(ctx, d, m)
}
