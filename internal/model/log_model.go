package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LogModel = (*customLogModel)(nil)

type (
	// LogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLogModel.
	LogModel interface {
		logModel
		FindTeamsByTeamID(ctx context.Context, teamID int64) ([]*Log, error)
		withSession(session sqlx.Session) LogModel
	}

	customLogModel struct {
		*defaultLogModel
	}
)

// NewLogModel returns a model for the database table.
func NewLogModel(conn sqlx.SqlConn) LogModel {
	return &customLogModel{
		defaultLogModel: newLogModel(conn),
	}
}

func (m *defaultLogModel) FindTeamsByTeamID(ctx context.Context, teamID int64) ([]*Log, error) {
	query := fmt.Sprintf("select %s from %s where `team_id` = ?", logRows, m.table)
	var resp []*Log
	err := m.conn.QueryRowsCtx(ctx, &resp, query, teamID)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customLogModel) withSession(session sqlx.Session) LogModel {
	return NewLogModel(sqlx.NewSqlConnFromSession(session))
}
