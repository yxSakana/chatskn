package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"chatskn/app/message/model"
	"chatskn/app/message/rpc/internal/svc"
	"chatskn/app/message/rpc/pb"
)

type StorageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStorageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StorageLogic {
	return &StorageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StorageLogic) Storage(in *pb.StorageReq) (*pb.StorageResp, error) {
	ret, err := l.svcCtx.MessageModel.Insert(l.ctx, &model.Message{
		SenderId:   in.Msg.SenderId,
		ReceiverId: in.Msg.ReceiverId,
		ChannelId:  in.Msg.ChannelId,
		Type:       in.Msg.Type,
		Content:    in.Msg.Content,
	})
	if err != nil {
		return nil, err
	}

	uid, err := ret.LastInsertId()

	return &pb.StorageResp{Id: uid}, nil
}
