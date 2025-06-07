package mqs

import (
	"context"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"

	"chatskn/app/ws/rpc/internal/config"
	"chatskn/app/ws/rpc/internal/svc"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqConsumerConf, NewBroadcastConsumer(ctx, svcContext)),
	}
}
