package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	RedisConf      redis.RedisConf
	GuildRpcConf   zrpc.RpcClientConf
	MessageRpcConf zrpc.RpcClientConf
	KqConsumerConf kq.KqConf
	KqPusherConf   struct {
		Brokers []string
		Topic   string
	}
}
