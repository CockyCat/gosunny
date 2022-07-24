package main

import (
	"ayzy_platform/job/internal/config"
	"ayzy_platform/job/internal/logic"
	"ayzy_platform/job/internal/svc"

	"context"
	"flag"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/job.yaml", "Specify the config file")

//Job服务主入口
func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	//
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	//logx.DisableStat()

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	//启动分布式任务服务
	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
