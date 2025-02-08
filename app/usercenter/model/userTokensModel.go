package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserTokensModel = (*customUserTokensModel)(nil)

type (
	// UserTokensModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTokensModel.
	UserTokensModel interface {
		userTokensModel
	}

	customUserTokensModel struct {
		*defaultUserTokensModel
	}
)

// NewUserTokensModel returns a model for the database table.
func NewUserTokensModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserTokensModel {
	return &customUserTokensModel{
		defaultUserTokensModel: newUserTokensModel(conn, c, opts...),
	}
}
