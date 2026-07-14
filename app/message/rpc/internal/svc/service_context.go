package svc

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"chatskn/app/message/model"
	"chatskn/app/message/rpc/internal/config"
	"chatskn/app/message/rpc/internal/esc"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel model.MessageModel
	Searcher     esc.MessageSearcher
}

func NewServiceContext(c config.Config) *ServiceContext {
	es, err := elasticsearch.NewTypedClient(c.ElasticSearchConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:       c,
		MessageModel: model.NewMessageModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		Searcher:     esc.NewEsClient(es),
	}
}
