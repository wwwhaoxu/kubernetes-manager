package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"kubernetes-manager/internal/pkg/log"
)

type K8sConfig struct {
	Name  string // 集群名称
	Host  string // 集群地址 https://xxxx:8443
	Token string // Token
	CA    string // CA证书
}

var (
	// to-do read from configuration file
	testconfig K8sConfig = K8sConfig{
		Name:  "test",
		Host:  "https://127.0.0.1:55608",
		Token: "eyJhbGciOiJSUzI1NiIsImtpZCI6IjNQcjdZTnBlcnNrME1PX2Vnb1V1X1F3X09FMXYwZERSUl9aZXQyWUgwN2MifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJla3MtYWRtaW4iLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC5uYW1lIjoiZWtzLWFkbWluIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQudWlkIjoiYjlkNWNhZWEtODAxZi00NTAwLWIzMjUtYjI1ZTdhNTZmOWQ5Iiwic3ViIjoic3lzdGVtOnNlcnZpY2VhY2NvdW50Omt1YmUtc3lzdGVtOmVrcy1hZG1pbiJ9.Ei0gBzGw1GzIELY3qwlbxmbG1hMYWZ2qhmW7qm18S0cvS_s_0u6vzhOXbxmX6pnO8r_4hBxF6s1KN9HGwzQcBup9xqDb83L6Q6aS-jKHkFqscXZrOGKU3Mt3xOXX10eJNDi1L8HCFE5OafnvO2hFmjqZGwxls5_e-aKSYv9eEWzfxpU9pQN_1hianPSCW0RyPbL4dC05U5OoGCjAvTc04jpfYBr6Ywu_jNWKdsdPCs3MYd7sc__geqaSAsBh8SgbKwft1XuZOSZy75zqgapFJ83xUXoFeig_imciwhSemRgdaqkmXZKYi9NQMgbPVsTQlD0MOpiL39556TSeszliug", // 存放 serviceaccount对应secret的token

	}
	prodconfig K8sConfig = K8sConfig{
		Name:  "prod",
		Host:  "https://172.16.20.10:6443",
		Token: "xxxxxxxx",
	}
	// 可继续添加其他集群配置
)

func NewK8sConfig(env string) (*kubernetes.Clientset, error) {

	var c K8sConfig
	switch env { // 多集群支持
	case "test":
		c = testconfig
	case "prod":
		c = prodconfig
	default:
		log.Fatalw("env not support", env)
	}

	config := &rest.Config{
		Host:            c.Host,
		BearerToken:     c.Token,
		BearerTokenFile: "",
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	return kubernetes.NewForConfig(config)

}
