package logic

import (
	"ayzy_platform/job/desc"
	"ayzy_platform/job/internal/svc"
	"context"

	"github.com/hibiken/asynq"
)

/**
* Description
* 作业调度注册中心
 */
type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//注册器，用于注册Job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//注册同步流向作业
	mux.Handle(desc.ScheduleSyncHs, NewSyncHsfromDFMHandler(l.svcCtx))

	return mux
}
