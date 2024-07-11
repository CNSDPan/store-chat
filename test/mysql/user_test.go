package mysql

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"store-chat/model/mysqls"
	"store-chat/tools/tools"
	"strconv"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	users := []mysqls.Users{}
	for i := 1; i < 6; i++ {
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
		users = append(users, mysqls.Users{
			UserID:    uuid,
			Token:     string(output),
			Status:    mysqls.USER_STATUS_1,
			Name:      "用户" + strconv.Itoa(i),
			Fund:      tools.EnterExchange(1000),
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	fmt.Printf("%v \r\n", users)
	res := mysqls.NewUserMgr().Create(&users)
	fmt.Printf("创建 %v", res.Error)
}
