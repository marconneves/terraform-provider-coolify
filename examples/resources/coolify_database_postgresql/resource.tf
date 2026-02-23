resource "coolify_database_postgresql" "test" {
  name             = "my-postgres"
  server_uuid      = "server-uuid"
  project_uuid     = "project-uuid"
  environment_name = "production"
  postgres_user    = "user"
  postgres_password = "password"
  postgres_db      = "database"
}
