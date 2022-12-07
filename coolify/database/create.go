package database

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/api/client"
)



func databaseCreateItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	// // 1. New Database
	// // 2. Set Tipe of DB
	// // 4. Set destination of DB
	// // 5. Update Base of configs DB
	// 6. Deploy DB
	// 7. Set is Public when is Public

	id, err := apiClient.NewDatabase()
	if err != nil {
		return err
	}
	d.SetId(*id)

	err = apiClient.SetEngineDatabase(*id, d.Get("engine").(string))
	if err != nil {
		return err
	}

	err = apiClient.SetDestinationDatabase(*id, d.Get("destinationId").(string))
	if err != nil {
		return err
	}

	databaseToUpdate := &client.UpdateDatabaseDTO{
		Name: d.Get("name").(string),
		Version: d.Get("engine_version").(string),
	}

	err = apiClient.UpdateDatabase(*id, databaseToUpdate)
	if err != nil {
		return err
	}


	return nil
}