data "coolify_team" "example" {
  id = "12345"
}

output "team_name" {
  value = data.coolify_team.example.name
}