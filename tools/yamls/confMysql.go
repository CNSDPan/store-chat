package yamls

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

type MysqlConf struct {
	Name       string      `json:",optional"`
	Separation int         `json:",optional"`
	MasterDB   string      `json:",optional"`         // 主
	SlaveDB    SlaveDBConf `json:",optional,inherit"` // 从
	AloneDB    string      `json:",optional"`         // 单机
}

// SlaveDBConf 从配置
type SlaveDBConf struct {
	Tag     []string `json:",optional"`
	Connect []string `json:",optional"`
}

var MysqlCon *MysqlConf

const (
	SEPARATION_YES = 1
	SEPARATION_NO  = 2
)

var separationName = map[int]string{
	SEPARATION_YES: "读写分离",
	SEPARATION_NO:  "单机",
}

func init() {
	// 获取配置文件的路径
	realPath := getCurrentDir()
	mysqlFilePath := realPath + "/file-mysql.yaml"
	mysqlFile := flag.String("mysql-f", mysqlFilePath, "the mysql config file")

	var c MysqlConf
	conf.MustLoad(*mysqlFile, &c)
	MysqlCon = &c
}
