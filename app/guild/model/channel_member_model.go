package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChannelMemberModel = (*customChannelMemberModel)(nil)

type (
	channelMemberTransHandle func(context context.Context, session sqlx.Session) error

	// ChannelMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChannelMemberModel.
	ChannelMemberModel interface {
		channelMemberModel

		Trans(ctx context.Context, fn channelMemberTransHandle) error
	}

	customChannelMemberModel struct {
		*defaultChannelMemberModel
	}
)

// NewChannelMemberModel returns a model for the database table.
func NewChannelMemberModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ChannelMemberModel {
	return &customChannelMemberModel{
		defaultChannelMemberModel: newChannelMemberModel(conn, c, opts...),
	}
}

func (m *customChannelMemberModel) Trans(ctx context.Context, fn channelMemberTransHandle) error {
	return m.TransactCtx(ctx, func(context context.Context, session sqlx.Session) error {
		return fn(context, session)
	})
}
