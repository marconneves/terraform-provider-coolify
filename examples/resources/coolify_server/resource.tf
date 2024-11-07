
terraform {
  required_providers {
    coolify = {
      source = "marconneves/coolify"
      version = "4.4.0-beta.1"
    }
  }
}

provider "coolify" {
  address = "https://cloud.ferramenta.digital"
  token = "3|Tt6iybaozpfPZojdlK5gGGotoxK8Cn1u1xTUCe2T1e65744d"
}

resource "coolify_server" "test" {
    name             = "example-server"
    ip               = "192.168.1.100"
    port             = "22"
    user             = "root"
    private_key_uuid = "xso0ooc4o0w4cswcwws8gswg"
}
