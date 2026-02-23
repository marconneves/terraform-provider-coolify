resource "coolify_database_mariadb" "mariadb" {
  name             = "my-mariadb-public"
  server_uuid      = "server-uuid"
  project_uuid     = "project-uuid"
  environment_name = "production"
  
  mariadb_database = "coolify"
  mariadb_user     = "coolify"
  mariadb_password = "very-secure-password"
  
  is_public        = true
  public_port      = 3306
  
  instant_deploy   = true
  
  limits_memory    = "512M"
  limits_cpus      = "0.5"
}
