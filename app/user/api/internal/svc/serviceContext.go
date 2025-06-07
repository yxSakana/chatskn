package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"user/api/internal/config"
	"user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
