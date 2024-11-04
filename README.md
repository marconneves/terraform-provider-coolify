# Documentation

[Coolify - Doc](https://docs.coollabs.io/coolify)<br/>
[Terraform Provider - Doc](https://registry.terraform.io/providers/themarkwill/coolify/latest/docs)

# terraform-provider-coolify
Provider of Coolify for Terraform

```
go install
go mod tidy
```

## Example Usage
When you .

```terraform
terraform {
  required_providers {
    coolify = {
      source = "marconneves/coolify"
      version = "4.0.2"
    }
  }
}

provider "coolify" {
  address = "url of coolify"
  token = "token"
}

resource "coolify_team" "my_team" {
  name           = "example-team"
}
```