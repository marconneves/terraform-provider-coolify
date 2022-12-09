terraform {
  required_providers {
    coolify = {
      source = "themarkwill/coolify"
      version = "0.0.19"
    }
  }
}

provider "coolify" {
  address = "url of coolify"
  token = "token"
}

resource "coolify_database" "my_db" {
  name           = "outro-db"

  engine {
    image = "redis"
    version = "7.0"
  }

  settings {
    destination_id = "id-destination"
    is_public      = true
  }
}