package store

import (
	"context"
	"gorm.io/gorm"
	"kubernetes-manager/internal/pkg/model"
)

type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

type users struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

func (u users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error

}
