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
      source = "themarkwill/coolify"
      version = "0.1.3"
    }
  }
}

provider "coolify" {
  address = "url of coolify"
  token = "token"
}

resource "coolify_database" "my_db" {
  name           = "outro-db"
  engine         = "redis:7.0"

  settings {
    destination_id = "id-destination"
    is_public      = true
  }
}
```