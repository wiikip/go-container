provider "kubernetes" {
  host = var.KUBE_HOST

  client_certificate = var.KUBE_CLIENT_CERTIFICATE
  client_key = var.KUBE_CLIENT_KEY
  cluster_ca_certificate = var.KUBE_CA_CERTIFICATE
}
terraform {
  backend "remote" {
    organization = "wiikip-org"

    workspaces {
      name = "go-container"
    }
  }
}