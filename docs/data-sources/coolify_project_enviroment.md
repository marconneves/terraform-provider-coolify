---
subcategory: "Project"
page_title: "Coolify: coolify_project_environment"
description: |-
  The environment data source to manage environments within a project on Coolify.
---

## Data Source: `coolify_project_environment`

The `coolify_project_environment` data source allows you to retrieve information about an environment within a project by its name and project UUID.

### Example Usage

```hcl
data "coolify_project_environment" "example" {
  name         = "environment_name"
  project_uuid = "project_uuid"
}

output "environment_id" {
  value = data.coolify_project_environment.example.id
}

output "environment_description" {
  value = data.coolify_project_environment.example.description
}

output "environment_created_at" {
  value = data.coolify_project_environment.example.created_at
}

output "environment_updated_at" {
  value = data.coolify_project_environment.example.updated_at
}
```

### Input Parameters

- `name` (Required): The name of the environment. This is used to identify the environment within the project.
- `project_uuid` (Required): The unique identifier of the project. This is used to locate the project that contains the environment.

### Output Attributes

- `id` (Output): The identifier of the environment.
- `name` (Output): The name of the environment.
- `description` (Output): A description of the environment.
- `project_id` (Output): The identifier of the project associated with the environment.
- `created_at` (Output): The creation timestamp of the environment.
- `updated_at` (Output): The last update timestamp of the environment.
- `status` (Output): The status of the environment.