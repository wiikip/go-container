package kube



type Container struct {
	Image string
}

type GoContainer interface {
	Create() func()	
}

