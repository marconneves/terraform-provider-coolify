---
page_title: "destination Resource - terraform-provider-coolify"
subcategory: ""
description: |-
  Destinations define where to deploy your application, database, or service.
---

# Resource `coolify_destination`

-> Visit the [Destinations - Docs Coolify](https://docs.coollabs.io/coolify/destinations) to see details of Destinations.

Destinations define where to deploy your application, database, or service.

~> **Note:** Now is available only Local Docker.

# coolify_destination

## Example Usage
When your not define network, we generate a UUID for here.
```hcl
resource "coolify_destination" "my_network" {
  name = "Project Terraform"
}
```

## Example Usage - with network
```hcl
resource "coolify_destination" "my_network" {
  name      = "Project Terraform"
  network   = "unique-network"
}
```

## Example Usage - Create a destination and a database
```hcl
resource "coolify_destination" "my_network" {
  name      = "Project Terraform"
  network   = "unique-network"
}

resource "coolify_database" "my_db" {
  name           = "my-db"
  engine         = "postgresql:13.8.0"

  settings {
    destination_id   = coolify_destination.my_network.id
    is_public        = true
    default_database = "postgres"
    user             = "user"
    password         = "password"
    root_password    = "root-password"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` -
  (Required)
  Required. The resource name of the destination, is only for preview.

* `network` -
  (Optional)
  Optional. Used to create docker networks within the defined Docker Engine. (When not is defined, we generate a UUID for here.)

* `engine` -
  (Optional)
  Optional. This is the socket of your local docker. Default is /var/run/docker.sock. I recommend you not change this value.


## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

* `id` -
  The ID of the destination.

The `status` block supports:
* `network` -
    The network, when whe generate the UUID, return here.