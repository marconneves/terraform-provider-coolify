resource "coolify_database_mysql" "test" {
  name             = "my-mysql-public"
  server_uuid      = "server-uuid"
  project_uuid     = "project-uuid"
  environment_name = "production"
  
  mysql_user       = "coolify"
  mysql_password   = "very-secure-password"
  mysql_database   = "coolify"
  
  is_public        = true
  public_port      = 3306
  
  instant_deploy   = true
  
  limits_memory    = "512M"
  limits_cpus      = "0.5"
}
