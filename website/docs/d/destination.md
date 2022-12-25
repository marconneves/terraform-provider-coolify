---
subcategory: "Destination"
page_title: "Coolify: coolify_destination"
description: |-
  Destinations define where to deploy your application, database, or service.
---

# Resource `coolify_destination`

-> Visit the [Destinations - Docs Coolify](https://docs.coollabs.io/coolify/destinations) to see details of Destinations.

Destinations define where to deploy your application, database, or service.

~> **Note:** Now is available only Local Docker.

## coolify_destination

### Example Usage
When your not define network, we generate a UUID for here.
```hcl
data "coolify_destination" "network" {
  network = "Project Terraform"
}
```

## Argument Reference

The following arguments are supported:

* `network` -
  (Required)
  Required. The name of the network if is unique.

## Attributes Reference

See [coolify_destination](https://registry.terraform.io/providers/themarkwill/coolify/latest/docs/resources/destination) resource for details of all the available attributes.