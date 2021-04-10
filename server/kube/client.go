package kube

import (
	"context"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type KubeClient struct {
	ClientSet kubernetes.Clientset
}

func (client *KubeClient) CreatePod(pod *core.Pod) (*core.Pod, error){
	pod, err := client.ClientSet.CoreV1().Pods(pod.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return nil,err
	}
	return pod, nil
}

func (client *KubeClient) GetPods() (*core.PodList, error){
	pods, err := client.ClientSet.CoreV1().Pods("go-container-managed").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}