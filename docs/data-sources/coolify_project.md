---
subcategory: "Project"
page_title: "Coolify: coolify_project"
description: |-
  The order resource to manage a team on Coolify.
---

# Data Source: `coolify_project`

The `coolify_project` data source allows you to retrieve information about a team by its identifier or name.

## Example Usage

```hcl
data "coolify_project" "example" {
  id = "uuid"
}

output "team_name" {
  value = data.coolify_project.example.name
}

output "team_description" {
  value = data.coolify_project.example.description
}

output "environment_ids" {
  value = [for env in data.coolify_project.example.environments : env.id]
}

output "environment_names" {
  value = [for env in data.coolify_project.example.environments : env.name]
}
```

## Input Parameters

- `id` (Optional): The identifier of the team. If specified, the data source will retrieve the team data associated with this ID.
- `name` (Optional): The name of the team. If specified, the data source will retrieve the team data associated with this name.

## Output Attributes

- `id` (Output): The identifier of the team.
- `name` (Output): The name of the team.
- `description` (Output): A description of the team.
- `environments` (Output): A list of environments associated with the project, each containing:
  - `id` (Output): The identifier of the environment.
  - `name` (Output): The name of the environment.
  - `description` (Output): A description of the environment.
  - `project_id` (Output): The associated project identifier.
  - `created_at` (Output): The creation timestamp of the environment.
  - `updated_at` (Output): The last update timestamp of the environment.
