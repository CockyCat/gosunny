package svc

import (
	"ayzy_platform/job/internal/config"
	"ayzy_platform/service/hs/rpc/hsrpcservice"

	"github.com/hibiken/asynq"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server

	//依赖的流向RPC
	HsRpc hsrpcservice.HsRpcService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		//调用service/hs/rpc/hsrpcservice的NewHsRpcService来New出一个rpc Client
		HsRpc: hsrpcservice.NewHsRpcService(zrpc.MustNewClient(c.HsRpcConf)),
	}
}
