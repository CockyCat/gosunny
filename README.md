# GoSunny - Best Practices Of Microservice Base on Golang(Gozero, gORM, gRPC, Asynq) To build a middle-platform 
# GoSunny是基于Golang技术栈的微服务工程化的全套解决方案，本案例用于构建企业数据中台

随着内部系统的增多，各个系统之间割裂，数据不互通，多个系统不但重复开发、重复造轮子，还不易扩展。为解决以上问题亟待需要一套中台系统来聚合和满足各个业务、部门的功能需求和使用场景。满足这些场景的话需要几部分重要组件：
1. 通用的账户系统
2. 通用的角色和权限：适配各个子系统的用户
  1. 功能角色
  2. 数据角色
3. 支持不同的数据源：适配各个子系统的不同的数据库

为满足以上要求，采用前端和后端分离的开发方式，和DDD（领域驱动设计）模型进行业务拆分。由于各个业务进行了原子的拆分，既能解决数据库耦合的带来的无法扩展的问题，又可以满足以后高并发场景。同时，随着业务的发展，可单独针对有需要的服务进行水平扩容。服务之间彼此隔离，互不影响。

本案例主要使用Go微服务技术栈实现一个企业数据中台，采用前后端分离的方式进行开发。



### 前端技术栈：
+ Vue

### 后端技术栈：

+ Go zero
+ gORM
+ gRPC
+ asynq


## 架构图 
![](https://s3.bmp.ovh/imgs/2022/07/14/5c9d8bdb23ab787d.jpg)


## API网关 - API Gateway
### 基于Nginx+Lua方案的技术选型：
+ APISIX
+ KONG



## 分布式任务作业 - Job
采用Asynq + Redis实现任务队列的管理，主要有：

+ 任务队列
+ 任务服务器
+ 定时任务
+ Redis监控


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
