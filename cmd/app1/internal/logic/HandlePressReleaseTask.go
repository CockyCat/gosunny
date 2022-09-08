package logic

import (
	"xxxx/cmd/app1/internal/svc"
	"fmt"

)

type PressReleaseHandler struct {
	svcCtx *svc.ServiceContext
}

func NewPressReleaseHandler(svcCtx *svc.ServiceContext) *PressReleaseHandler {
	return &PressReleaseHandler{
		svcCtx: svcCtx,
	}
}

func (l *PressReleaseHandler) ProcessTask() error {
	fmt.Println("处理作业A=====>:")
	return nil
}
