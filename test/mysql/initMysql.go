package mysql

import (
	"github.com/bwmarrin/snowflake"
	"strconv"
)

var ServiceIdRpc = "199"
var DBModel *struct {
	Node *snowflake.Node
}

func init() {
	var err error
	var nodeId int64
	var node *snowflake.Node
	nodeId, err = strconv.ParseInt(ServiceIdRpc, 0, 64)
	if err != nil {
		panic("转换 64位整形失败 " + err.Error())
	}
	node, err = snowflake.NewNode(nodeId)
	if err != nil {
		panic("new 节点失败 " + err.Error())
	}
	DBModel = &struct {
		Node *snowflake.Node
	}{Node: node}
}
