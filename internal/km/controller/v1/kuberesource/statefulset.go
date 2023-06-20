package kuberesource

import (
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubernetes-manager/internal/pkg/core"
)

func (ctrl *KuberesourceController) Get(ctx *gin.Context) {

	deploymentClient := ctrl.client.AppsV1().Deployments(ctx.Param("namespace"))
	// 获取 deployment 对象

	deployment, err := deploymentClient.Get(ctx, "coredns", metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, deployment)

}
