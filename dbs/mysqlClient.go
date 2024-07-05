package dbs

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"store-chat/tools/yamls"
	"time"
)

// 单列数据库实例
var dbMysql = map[string]*gorm.DB{}

// 读写分离的数据库实例
var dbRWMysql = map[string]*gorm.DB{}

func init() {
	c := yamls.MysqlCon
	ctx := context.Background()
	logg := logx.WithContext(ctx)
	logg.Infof("%s mysql connect db init...", c.Name)
	switch c.Separation {
	case yamls.SEPARATION_YES:
		initRWDB(c, logg)
	case yamls.SEPARATION_NO:
		initDB(c)
	default:
		panic("Separation undulated ")
	}
	logg.Infof("%s mysql connect db init ok", c.Name)
}

// initDB
// @Auth：
// @Desc：单个数据库实例
// @Date：2024-04-15 17:15:45
// @param：c
func initDB(c *yamls.MysqlConf) {
	var dbConn *gorm.DB
	var err error
	dbConn, err = gorm.Open(mysql.New(mysql.Config{
		DSN: c.MasterDB,
	}), &gorm.Config{})
	if err != nil {
		panic("单个数据库实例初始化失败")
	}
	db, e := dbConn.DB()
	if e != nil {
		panic("单个数据库实例: 获取 数据库 实例失败")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour)
	dbMysql[c.Name] = dbConn
}

// initRWDB
// @Auth：
// @Desc：mysql主从模式下读写分离实例
// @Date：2024-04-17 18:09:26
// @param：c
// @param：logg
func initRWDB(c *yamls.MysqlConf, logg logx.Logger) {
	var dbConn *gorm.DB
	var err error
	mDSN := c.MasterDB
	dbConn, err = gorm.Open(mysql.New(mysql.Config{
		DSN: mDSN,
	}), &gorm.Config{})
	if err != nil {
		panic("单个数据库实例初始化失败")
	}
	// replicate 主从,只拥有读权限
	replicates := []gorm.Dialector{}
	for i, dsn := range c.SlaveDB.Connect {
		sConf := mysql.Config{
			DSN: dsn,
		}
		replicates = append(replicates, mysql.New(sConf))
		logg.Infof("Separation tag $s", c.SlaveDB.Tag[i])
	}
	dbConn.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.New(mysql.Config{DSN: mDSN})},
		Replicas: replicates,
		Policy:   dbresolver.RandomPolicy{},
	}).SetMaxIdleConns(10).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))

	dbRWMysql[c.Name] = dbConn

	// 还有更加高级用法参考gorm官方文档-高级主题--Resolver,可设置规定的表或struct切换不同的连接使用
	// 当Register()内填入映射的数据库结构体 | 表名时,查询user | roles 数据会通过 slave2 的连接池去查询,不会通过 slave1去查
	//dbConn.Use(dbresolver.Register(dbresolver.Config{
	//	Sources:  []gorm.Dialector{mysql.New(mysql.Config{DSN: mDSN})},
	//	Replicas: []gorm.Dialector{mysql.New(mysql.Config{DSN: c.SlaveDB.Connect[0]})},
	//	Policy:   dbresolver.RandomPolicy{},
	//}).Register(dbresolver.Config{
	//	Replicas: []gorm.Dialector{mysql.New(mysql.Config{DSN: c.SlaveDB.Connect[1]})},
	//}, &User{}, "roles"))
}

// GetReadDB
// @Auth：
// @Desc：获取连接池的DB
// @Date：2024-04-15 13:41:29
// @param：dbName
// @return：db | nil
func GetReadDB(dbName string) (db *gorm.DB) {
	fmt.Printf("dbName %v", dbName)
	var ok bool
	if yamls.MysqlCon.Separation == yamls.SEPARATION_YES {
		db, ok = dbRWMysql[dbName]
		if ok {
			return db
		} else {
			return nil
		}
	} else if yamls.MysqlCon.Separation == yamls.SEPARATION_NO {
		db, ok = dbMysql[dbName]
		if ok {
			return db
		} else {
			return nil
		}
	}
	return nil
}
