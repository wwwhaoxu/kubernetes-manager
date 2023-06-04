package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"kubernetes-manager/internal/pkg/core"
	"kubernetes-manager/internal/pkg/errno"
	"kubernetes-manager/internal/pkg/log"
	v1 "kubernetes-manager/pkg/api/km/v1"
	pb "kubernetes-manager/pkg/proto/km/v1"
	"time"
)

func (ctrl UserController) List(c *gin.Context) {
	log.C(c).Infow("List user function called")
	var r v1.ListUserRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}
	resp, err := ctrl.b.Users().List(c, r.Offset, r.Limit)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}

func (ctrl UserController) ListUser(ctx context.Context, r *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	log.C(ctx).Infow("ListUser function called")
	resp, err := ctrl.b.Users().List(ctx, int(*r.Offset), int(*r.Limit))
	if err != nil {
		return nil, err
	}

	users := make([]*pb.UserInfo, 0, len(resp.Users))

	for _, u := range resp.Users {
		createdAt, _ := time.Parse("2006-01-02 15:04:05", u.CreatedAt)
		updatedAt, _ := time.Parse("2006-01-02 15:04:05", u.UpdatedAt)
		users = append(users, &pb.UserInfo{
			Username:  u.Username,
			Nickname:  u.Nickname,
			Email:     u.Email,
			Phone:     u.Phone,
			PostCount: u.PostCount,
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		})
	}
	ret := &pb.ListUserResponse{
		TotalCount: resp.TotalCount,
		Users:      users,
	}
	return ret, nil
}
