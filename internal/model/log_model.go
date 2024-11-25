package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LogModel = (*customLogModel)(nil)

type (
	// LogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLogModel.
	LogModel interface {
		logModel
		FindOneByEncryptCode(ctx context.Context, encryptCode string) (*Log, error)
		withSession(session sqlx.Session) LogModel
	}

	customLogModel struct {
		*defaultLogModel
	}
)

func (m *customLogModel) FindOneByEncryptCode(ctx context.Context, encryptCode string) (*Log, error) {
	if m == nil || m.conn == nil {
		return nil, errors.New("model or database connection is nil")
	}

	var log *Log
	query := fmt.Sprintf("select %s from %s where `encrypt_code` = ?", logRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &log, query, encryptCode)
	if err != nil {
		return nil, err
	}

	if log == nil {
		return nil, ErrNotFound
	}

	return log, nil
}

// NewLogModel returns a model for the database table.
func NewLogModel(conn sqlx.SqlConn) LogModel {
	return &customLogModel{
		defaultLogModel: newLogModel(conn),
	}
}

func (m *customLogModel) withSession(session sqlx.Session) LogModel {
	return NewLogModel(sqlx.NewSqlConnFromSession(session))
}
