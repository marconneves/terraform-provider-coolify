resource "coolify_database_mysql" "test" {
  name             = "my-mysql"
  server_uuid      = "server-uuid"
  project_uuid     = "project-uuid"
  environment_name = "production"
  mysql_user       = "user"
  mysql_password   = "password"
  mysql_database   = "database"
}
