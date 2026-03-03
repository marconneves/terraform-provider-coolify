
terraform {
  required_providers {
    coolify = {
      source = "marconneves/coolify"
    }
  }
}

provider "coolify" {
  address = "http://35.230.191.170:8000"
  token   = "1|h0gkseZZ4fDQcNqIeLXl1kQcfr28Rog5teQfWpGh47e19c35"
}

resource "coolify_project" "this" {
  name        = "my-new-project"
  description = "A project created via Terraform"
}

resource "coolify_database_redis" "this" {
  name             = "my-redis-db"
  server_uuid      = "dooskco4gc8w8ckss4ko0840"
  project_uuid     = coolify_project.this.id
  environment_name = "production"
  description      = "Redis database for my-new-project"
  image            = "redis:latest"
  is_public        = false
  instant_deploy   = true

  limits_memory = "256m"
  limits_cpus   = "0.5"
}
