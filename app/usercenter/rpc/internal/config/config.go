package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	JwtAuth struct {
		AccessSecret string
		ShortExpire  int64
		LongExpire   int64
	}

	DataSource string
	Cache      cache.CacheConf
}
