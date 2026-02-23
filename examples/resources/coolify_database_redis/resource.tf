resource "coolify_database_redis" "test" {
  name             = "my-redis"
  server_uuid      = "server-uuid"
  project_uuid     = "project-uuid"
  environment_name = "production"
}
