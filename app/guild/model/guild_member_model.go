package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GuildMemberModel = (*customGuildMemberModel)(nil)

type (
	guildMemberTransHandle func(context context.Context, session sqlx.Session) error

	// GuildMemberModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGuildMemberModel.
	GuildMemberModel interface {
		guildMemberModel

		Trans(ctx context.Context, fn guildMemberTransHandle) error
	}

	customGuildMemberModel struct {
		*defaultGuildMemberModel
	}
)

// NewGuildMemberModel returns a model for the database table.
func NewGuildMemberModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GuildMemberModel {
	return &customGuildMemberModel{
		defaultGuildMemberModel: newGuildMemberModel(conn, c, opts...),
	}
}

func (m *customGuildMemberModel) Trans(ctx context.Context, fn guildMemberTransHandle) error {
	return m.TransactCtx(ctx, func(context context.Context, session sqlx.Session) error {
		return fn(context, session)
	})
}
