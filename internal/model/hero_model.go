package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HeroModel = (*customHeroModel)(nil)

type (
	// HeroModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHeroModel.
	HeroModel interface {
		heroModel
		FindGroupIsNotPick(ctx context.Context) ([]*Hero, error)
	}

	customHeroModel struct {
		*defaultHeroModel
	}
)

func (m *customHeroModel) FindGroupIsNotPick(ctx context.Context) ([]*Hero, error) {
	if m == nil || m.conn == nil {
		return nil, errors.New("model or database connection is nil")
	}

	var heroes []*Hero
	query := fmt.Sprintf("select %s from %s where is_pick = 0 limit 2 for update",
		heroRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &heroes, query)
	if err != nil {
		return nil, err
	}

	if len(heroes) < 2 {
		return nil, errors.New("not enough available heroes")
	}
	return heroes, nil
}

// NewHeroModel returns a model for the database table.
func NewHeroModel(conn sqlx.SqlConn) HeroModel {
	return &customHeroModel{
		defaultHeroModel: newHeroModel(conn),
	}
}
