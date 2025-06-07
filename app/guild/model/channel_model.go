package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChannelModel = (*customChannelModel)(nil)

type (
	channelTransHandle func(context context.Context, session sqlx.Session) error

	// ChannelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChannelModel.
	ChannelModel interface {
		channelModel

		Trans(ctx context.Context, fn channelTransHandle) error
	}

	customChannelModel struct {
		*defaultChannelModel
	}
)

// NewChannelModel returns a model for the database table.
func NewChannelModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ChannelModel {
	return &customChannelModel{
		defaultChannelModel: newChannelModel(conn, c, opts...),
	}
}

func (m *customChannelModel) Trans(ctx context.Context, fn channelTransHandle) error {
	return m.TransactCtx(ctx, func(context context.Context, session sqlx.Session) error {
		return fn(context, session)
	})
}
