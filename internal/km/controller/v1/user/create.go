package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"kubernetes-manager/internal/pkg/core"
	"kubernetes-manager/internal/pkg/errno"
	"kubernetes-manager/internal/pkg/log"
	v1 "kubernetes-manager/pkg/api/km/v1"
)

const defaultMethods = "(GET)|(POST)|(PUT)|(DELETE)"

// Create 创建一个新的用户.
func (ctrl *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")

	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := ctrl.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	if _, err := ctrl.a.AddNamedPolicy("p", r.Username, "/v1/users/"+r.Username, defaultMethods); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}
	core.WriteResponse(c, nil, nil)
}
