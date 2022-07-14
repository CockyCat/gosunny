package svc

import (
	"ayzy_platform/service/analysis/api/internal/config"
	"ayzy_platform/service/hs/rpc/hsrpcservice"
	"ayzy_platform/service/settle/rpc/settlerpcservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	HsRpc     hsrpcservice.HsRpcService
	SettleRpc settlerpcservice.SettleRpcService //注入结算RPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		HsRpc:     hsrpcservice.NewHsRpcService(zrpc.MustNewClient(c.HsRpc)),
		SettleRpc: settlerpcservice.NewSettleRpcService(zrpc.MustNewClient(c.SettleRpc)),
	}
}
