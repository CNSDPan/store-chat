package mysqls

import (
	"context"
	"store-chat/dbs"
	"store-chat/tools/yamls"
)

const (
	STORE_STATUS_1 int8 = 1
	STORE_STATUS_2 int8 = 2
)

var StoreStatusName = map[int8]string{
	STORE_STATUS_1: "启用",
	STORE_STATUS_2: "禁用",
}

type Stores struct {
	ID        uint32 `gorm:"primaryKey;column:id" json:"-"`      // 自增ID
	StoreID   int64  `gorm:"column:store_id" json:"storeId"`     // IID
	Status    int8   `gorm:"column:status" json:"status"`        // 1=启用 2=禁用
	Name      string `gorm:"column:name" json:"name"`            // 昵称
	CreatedAt string `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt string `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
}

type StoresApi struct {
	ID        uint32 `gorm:"primaryKey;column:id" json:"-"`         // 自增ID
	StoreID   int64  `gorm:"column:store_id" json:"storeId,string"` // IID
	Status    int8   `gorm:"column:status" json:"status,string"`    // 1=启用 2=禁用
	Name      string `gorm:"column:name" json:"name"`               // 昵称
	CreatedAt string `gorm:"column:created_at" json:"createdAt"`    // 创建时间
	UpdatedAt string `gorm:"column:updated_at" json:"updatedAt"`    // 更新时间
}

type StoreMgr struct {
	*_BaseMgr
}

func StoreTableName() string {
	return "users"
}

func NewStoreMgr() *StoreMgr {
	db := dbs.GetReadDB(yamls.MysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &StoreMgr{_BaseMgr: &_BaseMgr{DB: db.Table(StoreTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Reset 重置gorm会话
func (obj *StoreMgr) Reset() *StoreMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *StoreMgr) GetUser(users Stores) (result StoresApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Stores{}).Where(&users).Find(&result).Error
	return
}

// Gets 获取批量结果
func (obj *StoreMgr) Gets() (results []*StoresApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Stores{}).Find(&results).Error

	return
}
