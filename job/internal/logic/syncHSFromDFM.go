package logic

import (
	"ayzy_platform/job/internal/svc"
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

// CloseHomestayOrderHandler close no pay homestayOrder
type SyncHsfromDFMHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSyncHsfromDFMHandler(svcCtx *svc.ServiceContext) *SyncHsfromDFMHandler {
	return &SyncHsfromDFMHandler{
		svcCtx: svcCtx,
	}
}

//同步处理
func (l *SyncHsfromDFMHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	fmt.Println("开始同步")

	return nil
}
