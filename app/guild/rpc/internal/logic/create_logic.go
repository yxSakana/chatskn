package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"chatskn/app/guild/model"
	"chatskn/app/guild/rpc/internal/svc"
	"chatskn/app/guild/rpc/pb"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *pb.CreateReq) (*pb.CreateResp, error) {
	_, err := l.svcCtx.GuildModel.Insert(l.ctx, &model.Guild{
		OwnerId: in.OwnerId,
		Name:    in.Name,
	})

	return &pb.CreateResp{}, err
}
