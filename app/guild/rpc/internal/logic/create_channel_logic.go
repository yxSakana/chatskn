package logic

import (
	"context"

	"chatskn/app/guild/rpc/internal/svc"
	"chatskn/app/guild/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChannelLogic {
	return &CreateChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateChannelLogic) CreateChannel(in *pb.CreateChannelReq) (*pb.CreateChannelResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateChannelResp{}, nil
}
