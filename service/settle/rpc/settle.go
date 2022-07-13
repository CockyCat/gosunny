package main

import (
	"flag"
	"fmt"

	"ayzy_platform/common/database"
	"ayzy_platform/service/settle/rpc/internal/config"
	"ayzy_platform/service/settle/rpc/internal/server"
	"ayzy_platform/service/settle/rpc/internal/svc"
	"ayzy_platform/service/settle/rpc/types/settle"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/settle.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//集成加载gorm初始化数据库
	database.DataBase(c.Mysql.DataSource)

	ctx := svc.NewServiceContext(c)
	svr := server.NewSettleRpcServiceServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		settle.RegisterSettleRpcServiceServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
