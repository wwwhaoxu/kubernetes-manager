package user

import (
	"github.com/gin-gonic/gin"
	"kubernetes-manager/internal/pkg/core"
	"kubernetes-manager/internal/pkg/errno"
	"kubernetes-manager/internal/pkg/log"
	v1 "kubernetes-manager/pkg/api/km/v1"
)

func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}
