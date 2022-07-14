package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"ayzy_platform/service/analysis/api/internal/svc"
	"ayzy_platform/service/analysis/api/internal/types"
	"ayzy_platform/service/hs/rpc/hsrpcservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetcustombyprovinceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetcustombyprovinceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetcustombyprovinceLogic {
	return &GetcustombyprovinceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetcustombyprovinceLogic) Getcustombyprovince(req *types.ProvinceCustomReq) (resp *types.CustomSummaryReply, err error) {
	//获取流向ID
	hsIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("id")))
	hsId, err := hsIdNumber.Int64()
	fmt.Println(hsId)
	if err != nil {
		return nil, err
	}
	//hs.IdRequest为RPC服务端定义的请求参数
	_, err = l.svcCtx.HsRpc.GetHs(l.ctx, &hsrpcservice.IdRequest{
		Id: hsId,
	})
	if err != nil {
		return nil, err
	}

	return &types.CustomSummaryReply{
		Factoryname:                "aaaaa",
		Top100Sum:                  200,
		CurrentMonthNewCustomCount: 300,
		CurrentYearCustomCount:     10000,
	}, nil
}
