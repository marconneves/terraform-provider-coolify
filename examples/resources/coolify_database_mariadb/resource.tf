resource "coolify_database_mariadb" "mariadb" {
  name             = "my-mariadb"
  server_uuid      = "server-uuid"
  project_uuid     = "project-uuid"
  environment_name = "production"
  
  mariadb_database = "coolify"
  mariadb_user     = "coolify"
}
