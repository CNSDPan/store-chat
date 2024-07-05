package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	ServiceId    string `json:",optional"`
	ServerName   string `json:",optional"`
	BucketNumber uint   `json:",optional"`
}
