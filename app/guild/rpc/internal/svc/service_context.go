package svc

import (
	"chatskn/app/guild/model"
	"chatskn/app/guild/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	GuildModel model.GuildModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		GuildModel: model.NewGuildModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
