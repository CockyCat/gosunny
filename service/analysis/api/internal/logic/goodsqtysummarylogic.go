package logic

import (
	"context"

	"ayzy_platform/service/analysis/api/internal/svc"
	"ayzy_platform/service/analysis/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsqtysummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsqtysummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsqtysummaryLogic {
	return &GoodsqtysummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsqtysummaryLogic) Goodsqtysummary(req *types.CommonReq) (resp *types.GoodsqtySummaryReply, err error) {
	// todo: add your logic here and delete this line

	return
}
