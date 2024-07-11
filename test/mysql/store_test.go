package mysql

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"store-chat/dbs"
	"store-chat/tools/yamls"
	"testing"
	"time"
)

func TestCreateStore(t *testing.T) {
	db := dbs.GetReadDB(yamls.MysqlCon.Name)
	var stores []map[string]interface{}
	for i := 1; i < 4; i++ {
		cmd := exec.Command("ksuid")
		output, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		output, err = simplifiedchinese.GB18030.NewDecoder().Bytes(output)
		if err != nil {
			panic(err)
		}
		uuid := DBModel.Node.Generate().Int64()
		stores = append(stores, map[string]interface{}{
			"store_id":   uuid,
			"status":     1,
			"name":       "",
			"created_at": time.Now().Format("2006-01-02 15:04:05"),
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	fmt.Printf("%v \r\n", stores)
	res := db.Table("store").Create(stores)
	fmt.Printf("创建 %v", res.Error)
}
