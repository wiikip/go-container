resource "kubernetes_namespace" "go-container" {
  metadata {
    name = "go-container"
  }

}