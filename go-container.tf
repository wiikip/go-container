resource "kubernetes_deployment" "go-container-server" {
  metadata {
    name      = "go-container-server"
    namespace = "go-container"
    labels = {
      app = "server"
    }
  }
  spec {
    replicas = 3
    selector {
      match_labels = {
        app = "server"
      }
    }
    template {
      metadata {
        labels = {
          app = "server"
        }
      }
      spec {
        container {
          env {
            name  = "PORT"
            value = var.port
          }
          image = "wiikip/go-container:${var.tag}"
          name  = "go-container"
        }
      }
    }
  }
}

resource "kubernetes_service" "go-container-service" {
  metadata {
    name      = "go-container-service"
    namespace = "go-container"
  }
  spec {
    port {
      port        = 80
      target_port = 3000
      protocol    = "TCP"
      node_port   = 31384
    }
    selector = {
      app = kubernetes_deployment.go-container-server.metadata.0.labels.app
    }
    type = "NodePort"
  }
}


output "nodeport" {
  value = kubernetes_service.go-container-service.spec.0.port.0.node_port

}
