package database

import (
	"context"
	"fmt"
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

	uri := GetUrl(item)

	d.Set("uri", uri)

	status := make(map[string]interface{})
	status["host"] = GetHost(item)
	status["port"] = GetPort(item)

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

func GetHost(item *client.Database) string {
	if item.Database.Settings.IsPublic {
		if item.Database.DestinationDocker.RemoteEngine {
			return item.Database.DestinationDocker.RemoteIpAddress
		} else if item.Settings.IpV4 != nil {
			return *item.Settings.IpV4
		} else {
			return *item.Settings.IpV6
		}
	} else {
		return *&item.Database.Id
	}
}

func GetPort(item *client.Database) int {
	if item.Database.Settings.IsPublic {
		return *item.Database.PublicPort
	} else {
		return item.PrivatePort
	}
}

func GetUser(databaseUser *string) string {

	if databaseUser != nil {
		return *databaseUser + ":"
	}
	return ""
}

func GenerateDbDetails(item *client.Database) (string, string, string) {
	db := item.Database
	databaseDefault := db.DefaultDatabase
	databaseDbUser := db.User
	databaseDbUserPassword := db.Password

	if db.Type == "mongodb" || db.Type == "edgedb" {
		if db.Type == "mongodb" {
			databaseDefault = "?readPreference=primary&ssl=false"
		}
		databaseDbUser = db.RootUser
		databaseDbUserPassword = db.RootPassword
	} else if db.Type == "redis" {
		databaseDefault = ""
		databaseDbUser = ""
	}

	return databaseDefault, databaseDbUser, databaseDbUserPassword
}

func GetUrl(item *client.Database) string {
	databaseDefault, databaseDbUser, databaseDbUserPassword := GenerateDbDetails(item)
	databaseDbUser = GetUser(&databaseDbUser)
	host := GetHost(item)
	port := strconv.Itoa(GetPort(item))

	return fmt.Sprintf(
		"%s://%s%s@%s:%s/%s",
		item.Database.Type,
		databaseDbUser,
		databaseDbUserPassword,
		host,
		port,
		databaseDefault,
	)
}
