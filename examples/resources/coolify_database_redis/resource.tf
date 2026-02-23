resource "coolify_database_redis" "test" {
  name             = "my-redis-public"
  server_uuid      = "server-uuid"
  project_uuid     = "project-uuid"
  environment_name = "production"
  
  image            = "redis:alpine"
  redis_password   = "very-secure-password"
  
  is_public        = true
  public_port      = 6379
  
  instant_deploy   = true
  
  limits_memory            = "512M"
  limits_cpus              = "0.5"
}
