---
subcategory: "Private Key"
page_title: "Coolify: coolify_private_key"
description: |-
  The private key data source to retrieve information about a private key on Coolify.
---

# Data Source: `coolify_private_key`

The `coolify_private_key` data source allows you to retrieve information about a private key by its UUID.

## Example Usage

```hcl
data "coolify_private_key" "example" {
  uuid = "example-uuid"
}

output "private_key_name" {
  value = data.coolify_private_key.example.name
}

output "private_key_description" {
  value = data.coolify_private_key.example.description
}

output "private_key_content" {
  value = data.coolify_private_key.example.private_key
}
```

## Input Parameters

- `uuid` (Required): The unique identifier of the private key. The data source will retrieve the private key data associated with this UUID.

## Output Attributes

- `id` (Output): The identifier of the private key.
- `uuid` (Output): The unique identifier of the private key.
- `name` (Output): The name of the private key.
- `description` (Output): A description of the private key.
- `private_key` (Output): The content of the private key.
- `is_git_related` (Output): Indicates if the key is related to Git.
- `team_id` (Output): The team identifier associated with the key.
- `created_at` (Output): The creation timestamp of the private key.
- `updated_at` (Output): The last update timestamp of the private key.
```
