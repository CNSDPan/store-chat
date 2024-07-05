package yamls

import (
	"runtime"
	"strings"
)

// getCurrentDir
// @Desc：获取yaml配置目录路径
// @return：string
func getCurrentDir() string {
	_, fileName, _, _ := runtime.Caller(1)
	aPath := strings.Split(fileName, "/")
	dir := strings.Join(aPath[0:len(aPath)-1], "/")
	return dir
}
