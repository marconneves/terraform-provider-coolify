---
subcategory: "Application"
page_title: "Coolify: coolify_application"
description: |-
  The order resource to create a application on Coolify.
---

# Resource `coolify_application`

-> Visit the [Applications - Docs Coolify](https://docs.coollabs.io/coolify/applications/) to see options of application and implementations.

The order resource to create a application on Coolify.

## coolify_application

### Example Usage - basic PostgreSQL

```hcl
resource "coolify_application" "test_item" {
  name   = "second-app"
  domain = "https://second-app.s.coolify.io"

  template {
    build_pack  = "node"
    image       = "node:14"
    build_image = "node:14"

    settings {
      install_command = "npm install"
      start_command   = "npm start"
      auto_deploy     = false
    }

    env {
      key   = "BASE_PROJECT"
      value = "production"
    }

    env {
      key   = "BASE_URL"
      value = "https://front.s.coolify.io"
    }
  }

  repository {
    repository_id = 579493141
    repository    = "cool-sample/sample-nodejs"
    branch        = "main"
  }

  settings {
    destination_id = "id-of-destination"
    source_id      = "id-of-source-on-coolify"
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` -
  (Required)
  Required. The resource name of the application, is only for preview.

- `domain` -
  (Optional)
  Optional. The domain of the application is public when not is bot.

- `is_bot` -
  (Optional)
  Optional. The application is bot or not. Default is false.

The `template` block supports:
  - `build_pack` -
    (Required)
    Required. The build pack of the application.
  - `image` -
    (Required)
    Required. The image of the application.
  - `build_image` -
    (Required)
    Required. The build image of the application.
  
In `template` block, the `settings` block supports:
  - `install_command` -
    (Optional)
    Optional. The install command of the application.
  - `build_command` -
    (Optional)
    Optional. The build command of the application.
  - `start_command` -
    (Optional)
    Optional. The start command of the application.
  - `auto_deploy` -
    (Optional)
    Optional. The auto deploy of the application. Default is false.

In `template` block, support multiples `env`, the `env` block supports:
  - `key` -
    (Required)
    Required. The key of the environment variable.
  - `value` -
    (Required)
    Required. The value of the environment variable.
  - `is_build_env` -
    (Optional)
    Optional. The environment variable is build or not. Default is false.

The `repository` block supports:
  - `repository_id` -
    (Required)
    Required. The repository id of the application.
  - `repository` -
    (Required)
    Required. The repository of the application.
  - `branch` -
    (Required)
    Required. The branch of the application.
  - `commit_hash` -
    (Optional)
    Optional. The commit hash of the application.

The `settings` block supports:
  - `destination_id` -
    (Required)
    Required. The destination id of the application.
  - `source_id` -
    (Required)
    Required. The source id of the application.



## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `id` -
  The ID of the application.

The `status` block supports: