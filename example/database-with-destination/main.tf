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

resource "coolify_destination" "network" {
  name    = "Network por Terraform"
  network = "terraform-network"
}

resource "coolify_database" "my_db" {
  name           = "outro-db"
  engine         = "redis:7.0"

  settings {
    destination_id = coolify_destination.network.id
    is_public      = true
    password       = "123456"
  }
}