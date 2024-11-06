data "coolify_team_members" "example" {
  team_id = "12345"
}

output "team_member_names" {
  value = data.coolify_team_members.example.members[*].name
}