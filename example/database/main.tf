terraform {
  required_providers {
    coolify = {
      source = "themarkwill/coolify"
      version = "0.2.1"
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
    destination_id = "destination id"
    is_public      = true
    password       = "123456"
  }
}