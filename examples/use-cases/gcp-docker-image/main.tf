terraform {
  required_providers {
    coolify = {
      source = "marconneves/coolify"
    }
  }
}

provider "coolify" {
  address = "https://coolify.example.com"
  # token via COOLIFY_TOKEN env var
}

variable "project_uuid" {
  type = string
}

variable "server_uuid" {
  type = string
}

variable "environment_name" {
  type    = string
  default = "production"
}

variable "image_tag" {
  type    = string
  default = "latest"
}

variable "environment_variables" {
  type    = map(string)
  default = {}
}

# Deploys a service from an image hosted on GCP Artifact Registry. The Coolify
# server must already have credentials configured to pull from the registry.
resource "coolify_application_dockerimage" "service" {
  name             = "checkout-api"
  description      = "Checkout API"
  project_uuid     = var.project_uuid
  server_uuid      = var.server_uuid
  environment_name = var.environment_name

  docker_registry_image_name = "us-central1-docker.pkg.dev/my-gcp-project/services/checkout-api"
  docker_registry_image_tag  = var.image_tag

  ports_exposes = "8080"
  domains       = "https://checkout.example.com"

  instant_deploy         = true
  is_force_https_enabled = true

  health_check_enabled     = true
  health_check_path        = "/healthz"
  health_check_return_code = 200

  limits_memory = "512M"
  limits_cpus   = "0.5"

  environment_variables = var.environment_variables
}

# Sensitive env that should appear only once in the Coolify UI, set as a
# separate resource so we can flip the per-variable flags.
resource "coolify_application_env" "stripe_key" {
  application_uuid = coolify_application_dockerimage.service.id
  key              = "STRIPE_SECRET_KEY"
  value            = "sk_live_..."
  is_shown_once    = true
  is_literal       = true
}
