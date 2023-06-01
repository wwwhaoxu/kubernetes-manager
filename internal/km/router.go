package km

import (
	"github.com/gin-gonic/gin"
	"kubernetes-manager/internal/km/controller/v1/user"
	"kubernetes-manager/internal/km/store"
	mw "kubernetes-manager/internal/pkg/middleware"

	"kubernetes-manager/internal/pkg/core"
	"kubernetes-manager/internal/pkg/errno"
	"kubernetes-manager/internal/pkg/log"
)

func installRoutes(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S)
	g.POST("/login", uc.Login)

	// 创建路由v1分组
	v1 := g.Group("v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mw.Authn())
		}
	}
	return nil
}
