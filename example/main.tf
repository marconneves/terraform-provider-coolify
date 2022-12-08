terraform {
  required_providers {
    coolify = {
      source = "themarkwill/coolify"
      version = "~>0.0.17"
    }
  }
}

provider "coolify" {
  address = "https://localhost:3300"
  token = "TOKEN OF COOLIFY"
}

resource "coolify_database" "my_db" {
  name           = "my-db"
  engine         = "postgresql"
  engine_version = "13.8.0"
  destination_id = "ID OF COOLIFY"
  is_public      = true
}