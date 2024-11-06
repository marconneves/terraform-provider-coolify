data "coolify_server" "example" {
  uuid = "example-id"
}

output "server_name" {
  value = data.coolify_server.example.name
}