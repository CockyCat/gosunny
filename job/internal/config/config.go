package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	Redis redis.RedisConf
	//注册流向和结算RPC到注册文件
	SettleRpcConf zrpc.RpcClientConf
	HsRpcConf     zrpc.RpcClientConf
}
