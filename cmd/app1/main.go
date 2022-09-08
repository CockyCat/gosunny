package main

import (
	"xxxx/cmd/app1/internal/config"
	"xxxx/cmd/app1/internal/logic"
	"xxxx/cmd/app1/internal/svc"

	"context"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/app1.yaml", "Specify the config file")

//Job服务主入口
func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, svcContext)

	//任务分发

	cronJob.HandlePressReleaseTask()
}
