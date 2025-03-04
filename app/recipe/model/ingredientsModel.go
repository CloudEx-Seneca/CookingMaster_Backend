package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IngredientsModel = (*customIngredientsModel)(nil)

type (
	// IngredientsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIngredientsModel.
	IngredientsModel interface {
		ingredientsModel
	}

	customIngredientsModel struct {
		*defaultIngredientsModel
	}
)

// NewIngredientsModel returns a model for the database table.
func NewIngredientsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) IngredientsModel {
	return &customIngredientsModel{
		defaultIngredientsModel: newIngredientsModel(conn, c, opts...),
	}
}
