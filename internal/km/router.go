package km

import (
	"github.com/gin-gonic/gin"
	"kubernetes-manager/internal/km/controller/v1/user"
	"kubernetes-manager/internal/km/store"
	mw "kubernetes-manager/internal/pkg/middleware"
	"kubernetes-manager/pkg/auth"

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

	authz, err := auth.NewAuthz(store.S.DB())

	// add policy
	if hasPolicy := authz.HasPolicy("admin", "/v1/users", user.DefaultMethods); !hasPolicy {
		authz.AddPolicy("admin", "/v1/users", user.DefaultMethods)
	}

	if hasPolicy := authz.HasPolicy("devops", "/v1/users", user.ReadMethods); !hasPolicy {
		authz.AddPolicy("devops", "/v1/users/belma", user.ReadMethods)
	}

	if err != nil {
		return err
	}
	uc := user.New(store.S, authz)
	g.POST("/login", uc.Login)

	// 创建路由v1分组
	v1 := g.Group("v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mw.Authn(), mw.Authz(authz))
			userv1.GET(":name", uc.Get) // 获取用户详情
			userv1.GET("", uc.List)
		}
	}
	return nil
}
