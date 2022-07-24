package svc

import (
	"ayzy_platform/job/internal/config"
	"fmt"

	"github.com/hibiken/asynq"
)

func newAsynqServer(c config.Config) *asynq.Server {
	//初始化asynq服务
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
		},
	)
}
