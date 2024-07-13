package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"store-chat/api/internal/config"
	"store-chat/api/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	AutoToken rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		AutoToken: middleware.NewAutoTokenMiddleware().Handle,
	}
}
