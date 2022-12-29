---
subcategory: "Destination"
page_title: "Coolify: coolify_destination"
description: |-
  Destinations define where to deploy your application, database, or service.
---

# Resource `coolify_destination` (Data Source)

Destinations define where to deploy your application, database, or service.


## coolify_destination

### Example Usage
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