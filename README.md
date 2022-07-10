# gozero_gorm_best-practices
# gozero和gORM最佳实践
Go zero是一款非常好用的微服务框架，但遗憾的是go zero自带的sqlx sqlc非常难用，所以结合自己的使用习惯，集成了gorm进来，同时封装了一层DAO层用来做一些基础DB操作

## 手动集成gORM
修改文件 api/xx.go

```
package main

import (
	"flag"
	"fmt"

	"YourNameSpace/common/database"

	"YourNameSpace/service/hs/api/internal/config"
	"YourNameSpace/service/hs/api/internal/handler"
	"YourNameSpace/service/hs/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/xx.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//集成gorm初始化数据库
	database.DataBase(c.Mysql.DataSource)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
```



## 使用模板集成gORM
未改模板时，用goctl生成的api或rpc里是没有集成gorm的，可用template里的响应文件替换到HOME文件夹下的.goctl目录模板

然后执行：

```
goctl api go -api xxxx.api -dir .
```
