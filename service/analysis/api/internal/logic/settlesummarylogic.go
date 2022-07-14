package logic

import (
	"context"

	"ayzy_platform/service/analysis/api/internal/svc"
	"ayzy_platform/service/analysis/api/internal/types"
	"ayzy_platform/service/settle/rpc/settlerpcservice"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SettlesummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSettlesummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SettlesummaryLogic {
	return &SettlesummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SettlesummaryLogic) Settlesummary(req *types.SettleSummaryReq) (resp *types.SettleSummaryReply, err error) {
	if err != nil {
		return nil, err
	}
	//请求结算类型汇总RPC

	var settleSummaryByTypeReply []*types.SettleSummaryByTypeReply
	summarybytype, err := l.svcCtx.SettleRpc.GetSummaryBySettleType(l.ctx, &settlerpcservice.SettleTypeSummaryRequest{})
	if err != nil {
		return nil, err
	}
	copier.Copy(&settleSummaryByTypeReply, summarybytype.List)

	//结算总金额

	var settleSumReply types.SettleSumReply
	sum, err := l.svcCtx.SettleRpc.GetSum(l.ctx, &settlerpcservice.SettleCommonRequest{})
	if err != nil {
		return nil, err
	}
	copier.Copy(&settleSumReply, sum)

	return &types.SettleSummaryReply{
		SettleSummaryByType: settleSummaryByTypeReply,
		SettleSum:           settleSumReply,
	}, nil
}
