resource "kubernetes_role" "go-container-role" {
    metadata {
      name = "go-container-role"
    }
    rule{
        api_groups = [ "" ]
        resources = [ "pods" ]
        verbs = [ "create", "get", "watch" ]
    }
  
}

resource "kubernetes_role_binding" "go-container-rb" {
    metadata {
      name = "go-container-rb"
    }
    role_ref {
      api_group = "rbac.authorization.k8s.io"
      kind = "Role"
      name = kubernetes_role.go-container-role.metadata.0.name
    }
    subject {
      kind = "ServiceAccount"
      name = "default"
      namespace = "go-container"
    }
  
}