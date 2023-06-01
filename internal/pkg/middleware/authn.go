package middleware

import (
	"github.com/gin-gonic/gin"
	"kubernetes-manager/internal/pkg/core"
	"kubernetes-manager/internal/pkg/errno"
	"kubernetes-manager/internal/pkg/known"
	"kubernetes-manager/pkg/token"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析jwt token
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}
