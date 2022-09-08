package logic

import (
	"xxxx/cmd/app1/internal/svc"
	"context"
)

/**
* Description
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

func (l *CronJob) HandlePressReleaseTask() {
	prHandler := NewPressReleaseHandler(l.svcCtx)
	prHandler.ProcessTask()
}

func (l *CronJob) HandleIndicatorDataTask() {

}

func (l *CronJob) HandlePrSeriesTask() {

}
