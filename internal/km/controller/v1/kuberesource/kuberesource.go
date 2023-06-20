package kuberesource

import (
	"k8s.io/client-go/kubernetes"
	"sync"
)

var (
	K    *KuberesourceController
	once sync.Once
)

type KuberesourceController struct {
	client *kubernetes.Clientset
	err    error
}

func New(client *kubernetes.Clientset) *KuberesourceController {
	// 确保 S 只被初始化一次
	once.Do(func() {
		K = &KuberesourceController{client: client}
	})
	return K
}
