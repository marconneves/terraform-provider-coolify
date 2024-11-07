resource "coolify_server" "test" {
    name             = "example-server"
    ip               = "192.168.1.100"
    port             = "22"
    user             = "example-user"
    private_key_uuid = "example-key-uuid"
}