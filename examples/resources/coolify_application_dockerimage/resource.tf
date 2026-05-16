resource "coolify_application_dockerimage" "api" {
  name             = "checkout-api"
  description      = "Checkout API deployed from GCP Artifact Registry"
  server_uuid      = "server-uuid"
  project_uuid     = "project-uuid"
  environment_name = "production"

  docker_registry_image_name = "us-central1-docker.pkg.dev/my-gcp-project/services/checkout-api"
  docker_registry_image_tag  = "v1.4.2"

  ports_exposes  = "8080"
  domains        = "https://checkout.example.com"

  instant_deploy         = true
  is_force_https_enabled = true

  health_check_enabled     = true
  health_check_path        = "/healthz"
  health_check_return_code = 200

  limits_memory = "512M"
  limits_cpus   = "0.5"

  environment_variables = {
    LOG_LEVEL    = "info"
    DATABASE_URL = "postgres://user:pass@db.example.com:5432/checkout"
    FEATURE_X    = "true"
  }
}
