data "coolify_application_dockerimage" "existing" {
  id = "application-uuid"
}

output "image" {
  value = "${data.coolify_application_dockerimage.existing.docker_registry_image_name}:${data.coolify_application_dockerimage.existing.docker_registry_image_tag}"
}

output "status" {
  value = data.coolify_application_dockerimage.existing.status
}
