data "coolify_private_key" "example" {
  id = "12345"
}

output "private_key_name" {
  value = data.coolify_private_key.example.name
}