package database

import (
	"context"
	"strconv"
	"strings"
	"terraform-provider-coolify/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func databaseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	databaseId := d.Id()

	item, err := apiClient.GetDatabase(databaseId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return diag.Errorf("error finding Item with ID %v", databaseId)
		}
	}

	d.Set("name", item.Database.Name)
	d.Set("engine", item.Database.Type+":"+item.Database.Version)

	settings := make(map[string]interface{})
	settings["destination_id"] = item.Database.DestinationDockerId
	settings["is_public"] = item.Database.Settings.IsPublic
	settings["append_only"] = item.Database.Settings.AppendOnly
	settings["default_database"] = item.Database.DefaultDatabase
	settings["user"] = item.Database.User
	settings["password"] = item.Database.Password
	settings["root_user"] = item.Database.RootUser
	settings["root_password"] = item.Database.RootPassword
	d.Set("settings", []interface{}{settings})

	status := make(map[string]interface{})
	if item.Database.Settings.IsPublic {
		if item.Settings.IpV4 != nil {
			status["host"] = *item.Settings.IpV4
		} else {
			status["host"] = *item.Settings.IpV6
		}
		status["port"] = strconv.Itoa(*item.Database.PublicPort)
	} else {
		status["host"] = *&item.Database.Id
		status["port"] = strconv.Itoa(item.PrivatePort)
	}

	if *&item.Database.DefaultDatabase != "" {
		status["default_database"] = *&item.Database.DefaultDatabase
	}
	if *&item.Database.User != "" {
		status["user"] = *&item.Database.User
	}
	if *&item.Database.Password != "" {
		status["password"] = *&item.Database.Password
	}
	if *&item.Database.RootUser != "" {
		status["root_user"] = *&item.Database.RootUser
	}
	if *&item.Database.RootPassword != "" {
		status["root_password"] = *&item.Database.RootPassword
	}

	d.Set("status", status)

	return nil
}
