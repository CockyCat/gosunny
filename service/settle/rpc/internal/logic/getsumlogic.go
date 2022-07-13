package logic

import (
	"context"
	"log"

	"ayzy_platform/common/dao"
	"ayzy_platform/service/settle/rpc/internal/svc"
	"ayzy_platform/service/settle/rpc/types/settle"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type TestResp struct {
	Totalmoney float64 `json:"totalmoney"`
}

type AyzyCbsSettlement struct {
	Id    string `json:"id"`
	Total string `json:"total"`
}

func NewGetSumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSumLogic {
	return &GetSumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSumLogic) GetSum(in *settle.SettleCommonRequest) (*settle.SettleSumResponse, error) {
	sql := "select sum(total) as totalmoney from ayzy_cbs_settlement" +
		" where 1 = ? "
	var settleSumResponse *settle.SettleSumResponse
	//var testResp TestResp
	if err := dao.QueryBySQL(sql, &settleSumResponse, "1"); err != nil {
		log.Println("未查到结果", err)
	}
	log.Println("结果为：", settleSumResponse)
	//return historicalSaleTopSupply
	return settleSumResponse, nil
}
