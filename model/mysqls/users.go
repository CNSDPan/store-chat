package mysqls

import (
	"context"
	"store-chat/dbs"
	"store-chat/tools/yamls"
)

const (
	USER_STATUS_0 int8 = 0
	USER_STATUS_1 int8 = 1
	USER_STATUS_2 int8 = 2
)

var StatusName = map[int8]string{
	USER_STATUS_1: "启用",
	USER_STATUS_2: "禁用",
}

type Users struct {
	ID        uint32 `gorm:"primaryKey;column:id" json:"-"`
	UserID    int64  `gorm:"column:user_id" json:"userId"`       // 用户IID
	Token     string `gorm:"column:token" json:"token"`          // token
	Status    int8   `gorm:"column:status" json:"status"`        // 1=启用 2=禁用
	Name      string `gorm:"column:name" json:"name"`            // 昵称
	Fund      int64  `gorm:"column:fund" json:"fund"`            // 用户资金,入库*1000【1000 = 1元】
	CreatedAt string `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt string `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
}

type UserApi struct {
	ID        uint32 `gorm:"primaryKey;column:id" json:"-"`
	UserID    int64  `gorm:"column:user_id" json:"userId,string"`       // 用户IID
	Token     string `gorm:"column:token" json:"token,string"`          // token
	Status    int8   `gorm:"column:status" json:"status,string"`        // 1=启用 2=禁用
	Name      string `gorm:"column:name" json:"name,string"`            // 昵称
	Fund      int64  `gorm:"column:fund" json:"fund,string"`            // 用户资金,入库*1000【1000 = 1元】
	CreatedAt string `gorm:"column:created_at" json:"createdAt,string"` // 创建时间
	UpdatedAt string `gorm:"column:updated_at" json:"updatedAt,string"` // 更新时间
}

type UsersMgr struct {
	*_BaseMgr
}

func UsersTableName() string {
	return "users"
}

func NewUserMgr() *UsersMgr {
	db := dbs.GetReadDB(yamls.MysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table(UsersTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Reset 重置gorm会话
func (obj *UsersMgr) Reset() *UsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *UsersMgr) GetUser(users Users) (result UserApi, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(&users).Find(&result).Error
	return
}

// Gets 获取批量结果
func (obj *UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *UsersMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Users, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

func (obj *UsersMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

// WithStatus status获取 1=启用 0=禁用
func (obj *UsersMgr) WithStatus(status int8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}
