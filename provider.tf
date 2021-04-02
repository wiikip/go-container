provider "kubernetes" {
  host = "https://138.195.138.77:6443"

  client_certificate = var.client_certificate
  client_key = var.client_key
  cluster_ca_certificate = var.ca_certificate
}