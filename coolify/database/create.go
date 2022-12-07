package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/api/client"
)



func databaseCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	// // 1. New Database
	// // 2. Set Tipe of DB
	// // 4. Set destination of DB
	// // 5. Update Base of configs DB
	// 6. Deploy DB
	// 7. Set is Public when is Public

	id, err := apiClient.NewDatabase()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*id)

	err = apiClient.SetEngineDatabase(*id, d.Get("engine").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	err = apiClient.SetDestinationDatabase(*id, d.Get("destination_id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	databaseToUpdate := &client.UpdateDatabaseDTO{
		Name: d.Get("name").(string),
		Version: d.Get("engine_version").(string),
		DefaultDatabase: "fist-db",
		DbUser: "user",
		DbUserPassword: "password",
	}

	err = apiClient.UpdateDatabase(*id, databaseToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}
	
	return nil
}