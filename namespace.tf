resource "kubernetes_namespace" "go-container" {
  metadata {
    name = "go-container"
  }

}
resource "kubernetes_namespace" "go-container-managed" {
  metadata {
    name = "go-container-managed"
  }
}