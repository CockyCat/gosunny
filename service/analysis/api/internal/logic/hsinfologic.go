package logic

import (
	"context"

	"ayzy_platform/service/analysis/api/internal/svc"
	"ayzy_platform/service/analysis/api/internal/types"
	"ayzy_platform/service/hs/rpc/hsrpcservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type HsinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHsinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HsinfoLogic {
	return &HsinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HsinfoLogic) Hsinfo(req *types.HsInfoReq) (resp *types.HsInfoReply, err error) {
	//获取流向ID
	if err != nil {
		return nil, err
	}
	hs, err := l.svcCtx.HsRpc.GetHs(l.ctx, &hsrpcservice.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.HsInfoReply{
		Id:          hs.Id,
		Dataid:      hs.Dataid,
		Factoryname: hs.Factoryname,
	}, nil
}
