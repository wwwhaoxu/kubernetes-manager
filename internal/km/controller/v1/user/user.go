package user

import (
	"kubernetes-manager/internal/km/biz"
	"kubernetes-manager/internal/km/store"
	"kubernetes-manager/pkg/auth"
	pb "kubernetes-manager/pkg/proto/km/v1"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	a *auth.Authz
	b biz.IBiz
	pb.UnimplementedKmServer
}

// New 创建一个 user controller.
func New(ds store.IStore, a *auth.Authz) *UserController {
	return &UserController{a: a, b: biz.NewBiz(ds)}
}
