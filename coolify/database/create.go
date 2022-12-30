package database

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-coolify/api/client"
)

func databaseCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	db := &Database{}
	db.Name = d.Get("name").(string)

	engineParts := strings.Split(d.Get("engine").(string), ":")
	db.Engine.Image = engineParts[0]
	db.Engine.Version = engineParts[1]

	settings := d.Get("settings").([]interface{})
	for _, setting := range settings {
		i := setting.(map[string]interface{})

		db.Settings.AppendOnly = i["append_only"].(bool)
		db.Settings.DestinationId = i["destination_id"].(string)
		db.Settings.IsPublic = i["is_public"].(bool) == true

		db.Settings.DefaultDatabase = i["default_database"].(string)
		db.Settings.User = i["user"].(string)
		db.Settings.Password = i["password"].(string)
		db.Settings.RootUser = i["root_user"].(string)
		db.Settings.RootPassword = i["root_password"].(string)
	}

	if db.Engine.Image == "mongodb" {
		if db.Settings.RootUser == "" {
			return diag.Errorf("default_database is required for MongoDB")
		}
		if db.Settings.RootPassword == "" {
			return diag.Errorf("default_database is required for MongoDB")
		}
	} else if db.Engine.Image == "mysql" {
		if db.Settings.DefaultDatabase == "" {
			return diag.Errorf("default_database is required for MySQL")
		}
		if db.Settings.User == "" {
			return diag.Errorf("user is required for MySQL")
		}
		if db.Settings.Password == "" {
			return diag.Errorf("password is required for MySQL")
		}
		if db.Settings.RootUser == "" {
			return diag.Errorf("default_database is required for MySQL")
		}
		if db.Settings.RootPassword == "" {
			return diag.Errorf("default_database is required for MySQL")
		}
	} else if db.Engine.Image == "mariadb" {
		if db.Settings.DefaultDatabase == "" {
			return diag.Errorf("default_database is required for MariaDB")
		}
		if db.Settings.User == "" {
			return diag.Errorf("user is required for MariaDB")
		}
		if db.Settings.Password == "" {
			return diag.Errorf("password is required for MariaDB")
		}
		if db.Settings.RootUser == "" {
			return diag.Errorf("default_database is required for MariaDB")
		}
		if db.Settings.RootPassword == "" {
			return diag.Errorf("default_database is required for MariaDB")
		}
	} else if db.Engine.Image == "postgresql" {
		if db.Settings.DefaultDatabase == "" {
			return diag.Errorf("default_database is required for PostgreSQL")
		}
		if db.Settings.User == "" {
			return diag.Errorf("user is required for PostgreSQL")
		}
		if db.Settings.Password == "" {
			return diag.Errorf("password is required for PostgreSQL")
		}
		if db.Settings.RootPassword == "" {
			return diag.Errorf("default_database is required for PostgreSQL")
		}
	} else if db.Engine.Image == "redis" {
		if db.Settings.Password == "" {
			return diag.Errorf("password is required for Redis")
		}
	} else if db.Engine.Image == "couchdb" {
		if db.Settings.DefaultDatabase == "" {
			return diag.Errorf("default_database is required for CouchDB")
		}
		if db.Settings.User == "" {
			return diag.Errorf("user is required for CouchDB")
		}
		if db.Settings.Password == "" {
			return diag.Errorf("password is required for CouchDB")
		}
		if db.Settings.RootUser == "" {
			return diag.Errorf("default_database is required for CouchDB")
		}
		if db.Settings.RootPassword == "" {
			return diag.Errorf("default_database is required for CouchDB")
		}
	} else if db.Engine.Image == "edgedb" {
		if db.Settings.DefaultDatabase == "" {
			return diag.Errorf("default_database is required for EdgeDB")
		}
		if db.Settings.RootUser == "" {
			return diag.Errorf("default_database is required for EdgeDB")
		}
		if db.Settings.RootPassword == "" {
			return diag.Errorf("default_database is required for EdgeDB")
		}
	}

	apiClient := m.(*client.Client)

	id, err := apiClient.NewDatabase()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*id)

	err = apiClient.SetEngineDatabase(*id, db.Engine.Image)
	if err != nil {
		return diag.FromErr(err)
	}

	err = apiClient.SetDestinationDatabase(*id, db.Settings.DestinationId)
	if err != nil {
		return diag.FromErr(err)
	}

	databaseToUpdate := &client.UpdateDatabaseDTO{
		Name:             db.Name,
		Version:          db.Engine.Version,
		DefaultDatabase:  &db.Settings.DefaultDatabase,
		DbUser:           &db.Settings.User,
		DbUserPassword:   &db.Settings.Password,
		RootUser:         &db.Settings.RootUser,
		RootUserPassword: &db.Settings.RootPassword,
	}

	err = apiClient.UpdateDatabase(*id, databaseToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "Starting database...")
	err = apiClient.StartDatabase(*id)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Trace(ctx, "Data base started")

	settingsToUpdate := &client.UpdateSettingsDatabaseDTO{
		IsPublic: db.Settings.IsPublic,
	}
	_, err = apiClient.UpdateSettings(*id, settingsToUpdate)
	if err != nil {
		return diag.FromErr(err)
	}

	return databaseRead(ctx, d, m)
}
