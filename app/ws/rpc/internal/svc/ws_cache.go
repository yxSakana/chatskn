package svc

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type IWsCache interface {
	CacheOfflineMessage(ctx context.Context, msgId int64) error
}

type WsCache struct {
	rdb *redis.Redis
}

func NewWsCache(r *redis.Redis) *WsCache {
	return &WsCache{rdb: r}
}

func (c *WsCache) CacheOfflineMessage(ctx context.Context, msgId int64) error {
	_, err := c.rdb.LpushCtx(ctx, "", msgId)
	return err
}

type OfflineMessage struct {
}
