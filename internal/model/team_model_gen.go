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
	teamFieldNames          = builder.RawFieldNames(&Team{})
	teamRows                = strings.Join(teamFieldNames, ",")
	teamRowsExpectAutoSet   = strings.Join(stringx.Remove(teamFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	teamRowsWithPlaceHolder = strings.Join(stringx.Remove(teamFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLotteryTeamIdPrefix = "cache:lottery:team:id:"
)

type (
	teamModel interface {
		Insert(ctx context.Context, data *Team) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Team, error)
		Update(ctx context.Context, data *Team) error
		Delete(ctx context.Context, id int64) error
		FindOneByEncryptCode(ctx context.Context, encrypt_code string) (*Team, error)
	}

	defaultTeamModel struct {
		sqlc.CachedConn
		conn  sqlx.SqlConn
		table string
	}

	Team struct {
		Id          int64          `db:"id"`
		EncryptCode sql.NullString `db:"encrypt_code"` // 队伍秘钥
		PickContent sql.NullString `db:"pick_content"` // 抽取结果
		IsPicked    int64          `db:"is_picked"`    // 是否抽取过
		UpdateTime  sql.NullTime   `db:"update_time"`  // 更新时间
	}
)

func newTeamModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTeamModel {
	return &defaultTeamModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		conn:       conn,
		table:      "`team`",
	}
}

func (m *defaultTeamModel) Delete(ctx context.Context, id int64) error {
	lotteryTeamIdKey := fmt.Sprintf("%s%v", cacheLotteryTeamIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, lotteryTeamIdKey)
	return err
}

func (m *defaultTeamModel) FindOne(ctx context.Context, id int64) (*Team, error) {
	lotteryTeamIdKey := fmt.Sprintf("%s%v", cacheLotteryTeamIdPrefix, id)
	var resp Team
	err := m.QueryRowCtx(ctx, &resp, lotteryTeamIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", teamRows, m.table)
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

func (m *defaultTeamModel) FindOneByEncryptCode(ctx context.Context, encrypt_code string) (*Team, error) {
	return nil, nil
}

func (m *defaultTeamModel) Insert(ctx context.Context, data *Team) (sql.Result, error) {
	lotteryTeamIdKey := fmt.Sprintf("%s%v", cacheLotteryTeamIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, teamRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.EncryptCode, data.PickContent, data.IsPicked)
	}, lotteryTeamIdKey)
	return ret, err
}

func (m *defaultTeamModel) Update(ctx context.Context, data *Team) error {
	lotteryTeamIdKey := fmt.Sprintf("%s%v", cacheLotteryTeamIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, teamRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.EncryptCode, data.PickContent, data.IsPicked, data.Id)
	}, lotteryTeamIdKey)
	return err
}

func (m *defaultTeamModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLotteryTeamIdPrefix, primary)
}

func (m *defaultTeamModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", teamRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTeamModel) tableName() string {
	return m.table
}
