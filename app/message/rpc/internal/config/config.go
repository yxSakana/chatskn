package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache             cache.CacheConf
	ElasticSearchConf elasticsearch.Config
}
