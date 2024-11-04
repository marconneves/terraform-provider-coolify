---
subcategory: "Team"
page_title: "Coolify: coolify_team"
description: |-
  The order resource to manage a team on Coolify.
---

# Data Source: `coolify_team`

The `coolify_team` data source allows you to retrieve information about a team by its identifier or name.

## Example Usage

```hcl
data "coolify_team" "example" {
  id = 1
}

output "team_name" {
  value = data.coolify_team.example.name
}

output "team_description" {
  value = data.coolify_team.example.description
}
```

## Input Parameters

- `id` (Optional): The identifier of the team. If specified, the data source will retrieve the team data associated with this ID.
- `name` (Optional): The name of the team. If specified, the data source will retrieve the team data associated with this name.

## Output Attributes

- `id` (Output): The identifier of the team.
- `name` (Output): The name of the team.
- `description` (Output): A description of the team.