package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"kubernetes-manager/internal/pkg/core"
	"kubernetes-manager/internal/pkg/errno"
	"kubernetes-manager/internal/pkg/log"
	v1 "kubernetes-manager/pkg/api/km/v1"
)

// ChangePassword 用来修改指定用户的密码.
func (ctrl *UserController) ChangePassword(c *gin.Context) {
	log.C(c).Infow("Change password function called")

	var r v1.ChangePasswordRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Users().ChangePassword(c, c.Param("name"), &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
