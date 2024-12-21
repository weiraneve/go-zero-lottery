package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LogModel = (*customLogModel)(nil)

type (
	// LogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLogModel.
	LogModel interface {
		logModel
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

func (m *customLogModel) withSession(session sqlx.Session) LogModel {
	return NewLogModel(sqlx.NewSqlConnFromSession(session))
}
