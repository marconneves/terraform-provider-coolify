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
resource "coolify_database" "my_db" {
  name           = "my-db"
  engine         = "postgresql:13.8.0"

  settings {
    destination_id = "id-of-destination"
    is_public      = true
    default_database = "postgres"
    user = "user"
    password = "password"
    root_password = "root-password"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` -
  (Required)
  Required. The resource name of the database, is only for preview.

* `engine` -
  (Required)
  Required. Engine of db, options: MongoDB, MySQL, MariaDB, PostgreSQL, Redis, CouchDB or EdgeDB with version of image See [Posibles Versions](#possibles-versions).


The `settings` block supports:
* `destination_id` -
  (Required)
  Required. The destination id on coolify.
* `is_public` -
  (Optional)
  Optional. If the database is public or not.
* `default_database` -
    (Optional)
    Optional. The default database of the database. *(Only for MySQL, MariaDB, PostgreSQL or CouchDB)
* `user` - 
    (Optional)
    Optional. The user of the database. *(Only for MySQL, MariaDB, PostgreSQL or CouchDB)
* `password` -
    (Optional)
    Optional. The password of the database. *(Only for MySQL, MariaDB, PostgreSQL, Redis or CouchDB)
* `root_user` -
    (Optional)
    Optional. The root user of the database. *(Only for MongoDB, MySQL, MariaDB, CouchDB or EdgeDB)
* `root_password` -
    (Optional)
    Optional. The root password of the database. *(Only for MongoDB, MySQL, MariaDB, PostgreSQL, CouchDB or EdgeDB)

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

* `id` -
  The ID of the database.
* `uri` -
  The Connection string of database.

The `status` block supports:
* `host` -
    The host of the database.
* `port` -
    The port of the database.
* `default_database` -
    The default database of the database. *(Only for MySQL, MariaDB, PostgreSQL or CouchDB)
* `user` -
    The user of the database. *(Only for MySQL, MariaDB, PostgreSQL or CouchDB)
* `password` -
    The password of the database. *(Only for MySQL, MariaDB, PostgreSQL, Redis or CouchDB)
* `root_user` -
    The root user of the database. *(Only for MongoDB, MySQL, MariaDB, CouchDB or EdgeDB)
* `root_password` -
    The root password of the database. *(Only for MongoDB, MySQL, MariaDB, PostgreSQL, CouchDB or EdgeDB)

## Possibles Versions

List of versions supported.

- **mongodb:**
    - 4.2
    - 4.4
    - 5.0
- **mysql:**
    - 5.7
    - 8.0
- **mariadb:**
    - 10.2
    - 10.3
    - 10.4
    - 10.5
    - 10.6
    - 10.7
    - 10.8
- **postgresql:**
    - 10.22.0
    - 11.17.0
    - 12.12.0
    - 13.8.0
    - 14.5.0
- **redis:**
    - 5.0
    - 6.0
    - 6.2
    - 7.0
- **couchdb:**
    - 2.3.1
    - 3.1.2
    - 3.2.2
- **edgedb:**
    - 1.4
    - 2.0
    - 2.1
    - latest
