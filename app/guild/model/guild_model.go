package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GuildModel = (*customGuildModel)(nil)

type (
	guildTransHandle func(context context.Context, session sqlx.Session) error

	// GuildModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGuildModel.
	GuildModel interface {
		guildModel

		Trans(ctx context.Context, fn guildTransHandle) error
	}

	customGuildModel struct {
		*defaultGuildModel
	}
)

// NewGuildModel returns a model for the database table.
func NewGuildModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GuildModel {
	return &customGuildModel{
		defaultGuildModel: newGuildModel(conn, c, opts...),
	}
}

func (m *customGuildModel) Trans(ctx context.Context, fn guildTransHandle) error {
	return m.TransactCtx(ctx, func(context context.Context, session sqlx.Session) error {
		return fn(context, session)
	})
}
