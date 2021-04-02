variable "port" {
 type =  string
 default = "3000"
}
variable "client_certificate" {
    type = string
}

variable "client_key" {
    type = string
}

variable "ca_certificate" {
    type = string
}

variable "cluster_endpoint"{
    type = string
}