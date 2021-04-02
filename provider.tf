provider "kubernetes" {
  host = "https://138.195.138.77:6443"

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