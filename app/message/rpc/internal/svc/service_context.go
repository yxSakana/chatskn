package svc

import (
	"chatskn/app/message/model"
	"chatskn/app/message/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel model.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		MessageModel: model.NewMessageModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
