package logic

import (
	"context"

	"ayzy_platform/service/analysis/api/internal/svc"
	"ayzy_platform/service/analysis/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomsummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCustomsummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomsummaryLogic {
	return &CustomsummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CustomsummaryLogic) Customsummary(req *types.CommonReq) (resp *types.ProvinceCustomReply, err error) {
	// todo: add your logic here and delete this line
	return &types.ProvinceCustomReply{
		Customname: "customname",
		Count:      "test order",
	}, nil
}
