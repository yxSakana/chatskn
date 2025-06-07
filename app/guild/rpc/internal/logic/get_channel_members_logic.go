package logic

import (
	"context"

	"chatskn/app/guild/rpc/internal/svc"
	"chatskn/app/guild/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChannelMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelMembersLogic {
	return &GetChannelMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc create() returns();
func (l *GetChannelMembersLogic) GetChannelMembers(in *pb.GetChannelMembersReq) (*pb.GetChannelMembersResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetChannelMembersResp{}, nil
}
