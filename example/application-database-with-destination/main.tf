terraform {
  required_providers {
    coolify = {
      source  = "themarkwill/coolify"
      version = "0.3.0-beta.2"
    }
  }
}

provider "coolify" {
  address = "https://services.b4.run"
  token   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjbGJmb2JuOHEwMGFoOWRtbzgxY2k5bHl4IiwidGVhbUlkIjoiY2xiOXZrbGp3MDAwdm1vOWRlbHFyNjFtaSIsInBlcm1pc3Npb24iOiJhZG1pbiIsImlzQWRtaW4iOnRydWUsImlhdCI6MTY3MDU0NDI4OX0.mc08PJ0b0JDQbU4Gt-uOVrawL5TyMewwbGD1jQperl0"
}

resource "coolify_destination" "network" {
  name    = "Seccond Application Terraform"
  network = "second-network-bibi"
}

resource "coolify_database" "smy_db" {
  name   = "second-application-db"
  engine = "redis:7.0"

  settings {
    destination_id = coolify_destination.network.id
    is_public      = true
    password       = "123456"
  }
}

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
    branch        = "develop"
  }

  settings {
    destination_id = coolify_destination.network.id
    source_id      = "clb9y09gs000f9dmod69f7dce"
  }
}