---
subcategory: "Team"
page_title: "Coolify: coolify_team"
description: |-
  The order resource to manage a team on Coolify.
---

# Resource `coolify_team`

## coolify_team

### Example to get a team
```hcl
data "coolify_team" "my_team" {
  team_id           = "team id"
}
```

## Argument Reference

The following arguments are supported:

* `team_id` -
  The ID of the Team.