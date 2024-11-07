terraform {
  required_providers {
    coolify = {
      source = "marconneves/coolify"
      version = "4.3.3"
    }
  }
}


provider "coolify" {
  address = "https://cloud.ferramenta.digital"
  token = "3|Tt6iybaozpfPZojdlK5gGGotoxK8Cn1u1xTUCe2T1e65744d"
}


resource "coolify_project" "test" {
  name        = "teste-project"
  description = "Descricao base"
}