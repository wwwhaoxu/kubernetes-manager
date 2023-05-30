package biz

import (
	"kubernetes-manager/internal/km/store"

	"kubernetes-manager/internal/km/biz/user"
)

type IBiz interface {
	Users() user.UserBiz
}

type biz struct {
	ds store.IStore
}

var _ IBiz = (*biz)(nil)

func NewBiz(ds store.IStore) *biz {
	return &biz{
		ds,
	}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b biz) Users() user.UserBiz {
	return user.New(b.ds)
}
