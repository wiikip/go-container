variable "port" {
 type =  string
 default = "3000"
}
variable "KUBE_CLIENT_CERTIFICATE" {
    type = string
}

variable "KUBE_CLIENT_KEY" {
    type = string
}

variable "KUBE_CA_CERTIFICATE" {
    type = string
}

variable "KUBE_HOST" {
    type = string
}

variable "tag" {
    type = string
    default = "latest" 
}