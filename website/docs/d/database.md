---
subcategory: "Database"
page_title: "Coolify: coolify_database"
description: |-
  The order resource to create a database on Coolify.
---

# Resource `coolify_database`

-> Visit the [Databases - Docs Coolify](https://docs.coollabs.io/coolify/databases) to see options of database and implementations.

The order resource to create a database on Coolify.

## coolify_database

### Example Usage - basic PostgreSQL
```hcl
data "coolify_database" "my_db" {
  database_id           = "database id"
}
```

## Argument Reference

The following arguments are supported:

* `database_id` -
  The ID of the database.

## Attributes Reference

See [coolify_destination](https://registry.terraform.io/providers/themarkwill/coolify/latest/docs/resources/database) resource for details of all the available attributes.