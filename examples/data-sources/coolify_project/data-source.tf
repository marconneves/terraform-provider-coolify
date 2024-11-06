data "coolify_project" "example" {
  id = "12345"
}

output "project_name" {
  value = data.coolify_project.example.name
}