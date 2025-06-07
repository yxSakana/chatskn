package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	messageTransHandle func(context context.Context, session sqlx.Session) error

	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel

		Trans(ctx context.Context, fn messageTransHandle) error
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn, c, opts...),
	}
}

func (m *customMessageModel) Trans(ctx context.Context, fn messageTransHandle) error {
	return m.TransactCtx(ctx, func(context context.Context, session sqlx.Session) error {
		return fn(context, session)
	})
}
