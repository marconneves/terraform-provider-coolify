# Use coolify_application_env when you need per-variable flags like
# is_preview, is_literal, is_multiline or is_shown_once. For plain
# KEY = "value" entries, prefer the `environment_variables` map on
# coolify_application_dockerimage.

resource "coolify_application_env" "secret_key" {
  application_uuid = coolify_application_dockerimage.api.id
  key              = "SECRET_KEY"
  value            = "s3cr3t-value-revealed-once"
  is_shown_once    = true
  is_literal       = true
}

resource "coolify_application_env" "tls_cert" {
  application_uuid = coolify_application_dockerimage.api.id
  key              = "TLS_CERT_PEM"
  value            = file("${path.module}/cert.pem")
  is_multiline     = true
  is_literal       = true
}
