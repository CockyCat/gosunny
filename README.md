# gozero+gorm+grpc best-practices - To build a middle-platform 
# gozero+gORM+grpc+Vue 最佳实践 构建企业数据中台
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




## Docker部署

以下是在MacOS构建Linux服务器的镜像用例

如果是在Mac M1芯片下构建镜像的话，而服务器是Linux的话，记得指定CPU架构为amd64，否则会报：
[8] System error: exec format error


### 构建docker镜像
```
docker build -t xxx_service_api:v1.1.4 --platform=linux/amd64 . 
```

### 打包MacOS上的本地Docker镜像
```
docker save -o xxx_api_servicev1.1.4.tar xxx_service_api:v1.1.4
```
### 上传镜像至服务器



### 服务器导入Docker镜像
```
sudo docker load -i xxx_api_service.tar
```
### 确定是否导入成功
```
sudo docker images
```
### 启动docker服务
```
docker run -p 8881:8881 xxx_service_api:v1.1.4
```


## Docker部署Etcd

### 拉取etcd镜像
```
docker pull bitnami/etcd:latest
```

### 启动etcd
```
#docker run --env ALLOW_NONE_AUTHENTICATION=yes -p 2379:2379 bitnami/etcd:latest

docker run -it --rm \
    --network app-tier \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    bitnami/etcd:latest etcdctl --endpoints http://etcd-server:2379 put /message Hello
```
