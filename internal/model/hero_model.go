package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HeroModel = (*customHeroModel)(nil)

type (
	// HeroModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHeroModel.
	HeroModel interface {
		heroModel
	}

	customHeroModel struct {
		*defaultHeroModel
	}
)

// NewHeroModel returns a model for the database table.
func NewHeroModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HeroModel {
	return &customHeroModel{
		defaultHeroModel: newHeroModel(conn, c, opts...),
	}
}
