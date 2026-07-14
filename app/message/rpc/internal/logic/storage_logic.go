package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

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
	m := &model.Message{
		SenderId:   in.Msg.SenderId,
		ReceiverId: in.Msg.ReceiverId,
		ChannelId:  in.Msg.ChannelId,
		Type:       in.Msg.Type,
		Content:    in.Msg.Content,
	}
	if err := l.svcCtx.MessageModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		ret, err := l.svcCtx.MessageModel.Insert(l.ctx, m)
		if err != nil {
			return err
		}

		uid, err := ret.LastInsertId()
		if err != nil {
			return err
		}
		m.Id = uid

		return l.svcCtx.Searcher.Index(l.ctx, m)
	}); err != nil {
		return nil, err
	}

	return &pb.StorageResp{Id: m.Id}, nil
}
