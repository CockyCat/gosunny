package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type FredApiConf struct {
	ApiKey string
}

type Config struct {
	service.ServiceConf
	MacroBoxRpcConf zrpc.RpcClientConf
	FredApiConf     FredApiConf
}
