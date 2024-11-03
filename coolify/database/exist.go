package database

import (
	"strings"

	sdk "github.com/marconneves/coolify-sdk-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func databaseExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	apiClient := m.(*sdk.Client)
	databaseId := d.Id()

	_, err := apiClient.GetDatabase(databaseId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
