// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	heroFieldNames          = builder.RawFieldNames(&Hero{})
	heroRows                = strings.Join(heroFieldNames, ",")
	heroRowsExpectAutoSet   = strings.Join(stringx.Remove(heroFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	heroRowsWithPlaceHolder = strings.Join(stringx.Remove(heroFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLotteryHeroIdPrefix = "cache:lottery:hero:id:"
)

type (
	heroModel interface {
		Insert(ctx context.Context, data *Hero) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Hero, error)
		Update(ctx context.Context, data *Hero) error
		Delete(ctx context.Context, id int64) error
	}

	defaultHeroModel struct {
		sqlc.CachedConn
		table string
	}

	Hero struct {
		Id     int64  `db:"id"`
		Name   string `db:"name"`    // 英雄名
		Line   int64  `db:"line"`    // 英雄分路
		IsPick int64  `db:"is_pick"` // 是否被选择
	}
)

func newHeroModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultHeroModel {
	return &defaultHeroModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`hero`",
	}
}

func (m *defaultHeroModel) Delete(ctx context.Context, id int64) error {
	lotteryHeroIdKey := fmt.Sprintf("%s%v", cacheLotteryHeroIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, lotteryHeroIdKey)
	return err
}

func (m *defaultHeroModel) FindOne(ctx context.Context, id int64) (*Hero, error) {
	lotteryHeroIdKey := fmt.Sprintf("%s%v", cacheLotteryHeroIdPrefix, id)
	var resp Hero
	err := m.QueryRowCtx(ctx, &resp, lotteryHeroIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", heroRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultHeroModel) Insert(ctx context.Context, data *Hero) (sql.Result, error) {
	lotteryHeroIdKey := fmt.Sprintf("%s%v", cacheLotteryHeroIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, heroRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Line, data.IsPick)
	}, lotteryHeroIdKey)
	return ret, err
}

func (m *defaultHeroModel) Update(ctx context.Context, data *Hero) error {
	lotteryHeroIdKey := fmt.Sprintf("%s%v", cacheLotteryHeroIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, heroRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Line, data.IsPick, data.Id)
	}, lotteryHeroIdKey)
	return err
}

func (m *defaultHeroModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLotteryHeroIdPrefix, primary)
}

func (m *defaultHeroModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", heroRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultHeroModel) tableName() string {
	return m.table
}
