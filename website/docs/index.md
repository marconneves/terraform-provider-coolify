---
page_title: "Provider: Coolify"
subcategory: ""
description: |-
  Terraform provider for interacting with Coolify API.
---

# Coolify Provider

[Coolify](https://coolify.io) is a self-hostable, all-in-one solution to host your applications, databases, or other open-source services with a few simple clicks.

It's an alternative software to [Heroku](https://www.heroku.com/) and [Netlify](https://www.netlify.com/) and other alternatives out there.

You can try it out before installing it: [Live demo](https://demo.coolify.io/)

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
terraform {
  required_providers {
    coolify = {
      source = "themarkwill/coolify"
      version = "0.2.1"
    }
  }
}

provider "coolify" {
  address = "http://demo.coolify.io"
  token = "TOKEN"
}
```

# Schema

## Required

* `address` - (String, Optional) Coolify API address
* `token` - (String, Optional) Token of user authorized