package svc

import (
	"chatskn/app/ws/rpc/internal/webs"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"chatskn/app/guild/rpc/guild"
	"chatskn/app/message/rpc/message"
	"chatskn/app/ws/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Hub         *webs.Hub
	Cache       IWsCache
	KqPublisher *kq.Pusher
	GuildRpc    guild.GuildZrpcClient
	MessageRpc  message.MessageZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	r := redis.MustNewRedis(c.RedisConf)
	return &ServiceContext{
		Config:      c,
		Hub:         webs.NewHub(),
		Cache:       NewWsCache(r),
		KqPublisher: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic, kq.WithAllowAutoTopicCreation()),
		GuildRpc:    guild.NewGuildZrpcClient(zrpc.MustNewClient(c.GuildRpcConf)),
		MessageRpc:  message.NewMessageZrpcClient(zrpc.MustNewClient(c.MessageRpcConf)),
	}
}
