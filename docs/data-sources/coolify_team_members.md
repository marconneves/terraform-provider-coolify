---
subcategory: "Team"
page_title: "Coolify: coolify_team"
description: |-
  The order resource to manage a team on Coolify.
---

# Data Source: `coolify_team_members`

The `coolify_team_members` data source allows you to retrieve a list of members associated with a specific team.

## Example Usage

```hcl
data "coolify_team_members" "example" {
  team_id = 1
}

output "team_members" {
  value = data.coolify_team_members.example.members
}
```

## Input Parameters

- `team_id` (Required): The identifier of the team whose members you want to retrieve.

## Output Attributes

- `team_id` (Output): The identifier of the team.
- `members` (Output): A list of team members, where each member includes:
  - `id`: The identifier of the member.
  - `name`: The name of the member.
  - `email`: The email address of the member.