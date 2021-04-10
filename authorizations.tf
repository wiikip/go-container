resource "kubernetes_role" "go-container-role" {
  metadata {
    name      = "go-container-role"
    namespace = "go-container-managed"
  }
  rule {
    api_groups = [""]
    resources  = ["pods"]
    verbs      = ["create", "get", "watch", "list"]
  }

}

resource "kubernetes_role_binding" "go-container-rb" {
  metadata {
    name = "go-container-rb"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
    name      = kubernetes_role.go-container-role.metadata.0.name
    namespace = "go-container-managed"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "default"
    namespace = "go-container"
  }

}