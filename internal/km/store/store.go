package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	// 全局变量，方便其它包直接调用已初始化好的 S 实例.
	S *datastore
)

// IStore 定义了 Store 层需要实现的方法.
type IStore interface {
	DB() *gorm.DB
	Users() UserStore
}

// datastore 是 IStore 的一个具体实现.
type datastore struct {
	db *gorm.DB
}

var _ IStore = (*datastore)(nil)

func NewStore(db *gorm.DB) *datastore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datastore{db: db}
	})
	return S
}

func (ds *datastore) Users() UserStore {
	return newUsers(ds.db)

}

func (ds *datastore) DB() *gorm.DB {
	return ds.db
}
