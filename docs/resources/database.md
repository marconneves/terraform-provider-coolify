---
page_title: "database Resource - terraform-provider-coolify"
subcategory: ""
description: |-
  The order resource to create a database on Coolify.
---

# Resource `coolify_database`

-> Visit the [Databases - Docs Coolify](https://docs.coollabs.io/coolify/databases) to see options of database and implementations.

The order resource to create a database on Coolify.

## Example Usage

```terraform
resource "coolify_database" "my_db" {
  name           = "my-db"
  engine         = "postgresql"
  engine_version = "13.8.0"
  destination_id = "id_of_destination"

  settings {
    is_public = true
  }
}
```

## Argument Reference

- `name` - (Required) Name of project.
- `engine` - (Required) Engine of db, options: MongoDB, MySQL, MariaDB, PostgreSQL, Redis, CouchDB or EdgeDB.
- `engine_version` (Required) Version of engine. See [Posibles Versions](#possibles-versions) below for details.

### Possibles Versions

List of versions suported.
**postgresql:**
- 13.8.0

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.