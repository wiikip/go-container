resource "kubernetes_deployment" "go-container-server" {
    metadata {
      name="go-container-server"
      namespace="go-container"
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
    template{
      metadata{
        labels = {
          app = "server"
        }
      }
      spec{
        container{
          env {
            name= "PORT"
            value = var.port
          }
          image = "wiikip/go-container:latest"
          name = "go-container"
          }
        }
      }
    }
  }

resource "kubernetes_service" "go-container-service" {
    metadata {
      name = "go-container-service"
      namespace = "go-container"
    }
spec {
    port{
      port = 80
      target_port = 3000
      protocol = "TCP"
    }
    selector = {
      app = kubernetes_deployment.go-container-server.metadata.0.labels.app
    }
    type = "NodePort"
  }
}

resource "kubernetes_ingress" "go-container-ingress" {
    metadata {
      name = "go-container-ingress"
      namespace = "go-container"
    }
    spec {
    rule{
      host = "wiikip.viarezo.fr"
      http{
        path{
          path ="/"
          backend{
            service_name = kubernetes_service.go-container-service.metadata.0.name
            service_port = 80
          }
        }
      }

    }
  } 
  
}

output "nodeport" {
  value=kubernetes_service.go-container-service.spec.0.port.0.node_port
  
}
