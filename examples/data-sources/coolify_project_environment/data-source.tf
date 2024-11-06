data "coolify_project_environment" "example" {
  project_uuid = "uuid_project"
  name = "production"
}

output "project_name" {
  value = data.coolify_project_environment.example.name
}