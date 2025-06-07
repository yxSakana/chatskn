package mqs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"chatskn/app/guild/rpc/guild"
	"chatskn/app/message/model"
	"chatskn/app/message/rpc/message"
	"chatskn/app/ws/rpc/internal/svc"
)

type BroadcastConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBroadcastConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *BroadcastConsumer {
	return &BroadcastConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BroadcastConsumer) Consume(ctx context.Context, _, val string) error {
	// text -- 直接推送
	// image、文件 -- 上传到服务器，client通过链接获取
	logx.Infof("[BroadcastConsumer]: fetch message %s", val)
	hub := l.svcCtx.Hub
	msgByte := []byte(val)

	var msg model.Message
	if err := json.Unmarshal(msgByte, &msg); err != nil {
		return err
	}
	if msg.ReceiverId == 0 && msg.ChannelId == 0 {
		return fmt.Errorf("[BroadcastConsumer] 消息提取错误")
	}

	// -- save to db --
	msgRest, err := l.svcCtx.MessageRpc.Storage(l.ctx, &message.StorageReq{
		Msg: &message.Message{
			SenderId:   msg.SenderId,
			ReceiverId: msg.ReceiverId,
			ChannelId:  msg.ChannelId,
			Type:       msg.Type,
			Content:    msg.Content,
		},
	})
	if err != nil {
		return err
	}

	// -- 分发 ---
	isPrivate, isChannel := msg.ReceiverId != 0, msg.ChannelId != 0
	if isPrivate {
		// 私聊
		if client, ok := hub.Clients[msg.ReceiverId]; ok {
			client.Send <- msgByte
		}
	} else if isChannel {
		// 频道
		rest, err := l.svcCtx.GuildRpc.GetChannelMembers(l.ctx, &guild.GetChannelMembersReq{Id: msg.ChannelId})
		if err != nil {
			return err
		}
		for _, member := range rest.MembersInfo {
			if cli, ok := hub.Clients[member.UserId]; ok {
				cli.Send <- msgByte
			} else {
				_ = l.svcCtx.Cache.CacheOfflineMessage(l.ctx, msgRest.Id)
			}
		}
	}

	return nil
}
