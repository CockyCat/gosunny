package logic

import (
	"context"
	"log"

	"ayzy_platform/common/dao"
	"ayzy_platform/service/settle/rpc/internal/svc"
	"ayzy_platform/service/settle/rpc/types/settle"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSummaryBySettleTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSummaryBySettleTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSummaryBySettleTypeLogic {
	return &GetSummaryBySettleTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSummaryBySettleTypeLogic) GetSummaryBySettleType(in *settle.SettleTypeSummaryRequest) (*settle.SettleTypeSummaryListResponse, error) {
	// todo: add your logic here and delete this line
	sql := "select type as typename, sum(total) as totalmoney from ayzy_cbs_settlement" +
		" where 1 = ? " +
		" group by type "
	var settleTypeSummaryResponse []*settle.SettleTypeSummaryResponse
	if err := dao.QueryBySQL(sql, &settleTypeSummaryResponse, "1"); err != nil {
		log.Println("未查到结果", err)
	}
	lists := make([]*settle.SettleTypeSummaryResponse, 0)
	var typename string
	for _, s := range settleTypeSummaryResponse {

		switch s.Typename {
		case "4":
			typename = "设备"
		case "5":
			typename = "基础保养费"
		case "6":
			typename = "流向奖励费"
		case "7":
			typename = "会员积分费"
		case "8":
			typename = "流向信息费"
		case "9":
			typename = "设备确权"
		}
		settleResp := settle.SettleTypeSummaryResponse{
			Typename:   typename,
			Totalmoney: s.Totalmoney,
		}
		lists = append(lists, &settleResp)
	}
	return &settle.SettleTypeSummaryListResponse{
		List: lists,
	}, nil
}
