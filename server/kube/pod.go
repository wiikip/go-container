package kube

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BuildPayload struct {
	Name string `mapstructure:"name"`
	Uri  string `mapstructure:"uri"`
}

func NewPod(buildInfos BuildPayload) *core.Pod {
	return &core.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      buildInfos.Name,
			Namespace: "go-container-managed",
			Labels: map[string]string{
				"app": "demo",
			},
		},
		Spec: core.PodSpec{
			Containers: []core.Container{
				{
					Name:            buildInfos.Uri,
					Image:           buildInfos.Uri,
					ImagePullPolicy: core.PullIfNotPresent,
					Command: []string{
						"sleep",
						"3600",
					},
				},
			},
		},
	}

}
