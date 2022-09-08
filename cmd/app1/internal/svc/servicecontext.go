package svc

import (
	"xxxx/cmd/app1/internal/config"
	"xxxx/service/app1/rpc/macrobox"


	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	
	App1Rpc app1.App1Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		App1Service: app1.NewApp1(zrpc.MustNewClient(c.App1RpcConf)),
	}
}
