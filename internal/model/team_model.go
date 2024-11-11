package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TeamModel = (*customTeamModel)(nil)

type (
	// TeamModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTeamModel.
	TeamModel interface {
		teamModel
		FindOneByEncryptCode(ctx context.Context, encrypt_code string) (*Team, error)
	}

	customTeamModel struct {
		*defaultTeamModel
	}
)

func (m *defaultTeamModel) FindOneByEncryptCode(ctx context.Context, encrypt_code string) (*Team, error) {
	query := fmt.Sprintf("select %s from %s where `encrypt_code` = ?", teamRows, m.table)
	var resp Team
	err := m.conn.QueryRowCtx(ctx, &resp, query, encrypt_code)
	switch {
	case err == sqlx.ErrNotFound:
		return nil, ErrNotFound
	case err != nil:
		return nil, err
	default:
		return &resp, nil
	}
}

// NewTeamModel returns a model for the database table.
func NewTeamModel(conn sqlx.SqlConn) TeamModel {
	return &customTeamModel{
		defaultTeamModel: newTeamModel(conn),
	}
}
