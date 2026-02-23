resource "coolify_database_postgresql" "test" {
  name             = "my-postgres-public"
  server_uuid      = "server-uuid"
  project_uuid     = "project-uuid"
  environment_name = "production"
  
  postgres_user     = "coolify"
  postgres_password = "very-secure-password"
  postgres_db       = "coolify"
  
  is_public        = true
  public_port      = 5432
  
  instant_deploy   = true
  
  limits_memory    = "512M"
  limits_cpus      = "0.5"
}
