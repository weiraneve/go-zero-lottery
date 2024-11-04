package logic

import (
	"context"
	"errors"
	"strings"

	"lottery/internal/svc"
	"lottery/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PickLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Pick a team by encryptCode
func NewPickLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PickLogic {
	return &PickLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PickLogic) Pick(req *types.PickRequest) (resp *types.PickResponse, err error) {
	team, err := l.svcCtx.TeamModel.FindOneByEncryptCode(l.ctx, req.EncryptCode)
	if err != nil {
		return nil, err
	}

	if team == nil {
		return nil, errors.New("team not found, please check your encryptCode")
	}

	heroes, err := l.svcCtx.HeroModel.FindGroupIsNotPick(l.ctx)
	if err != nil {
		return nil, err
	}

	names := make([]string, len(heroes))
	for i, hero := range heroes {
		names[i] = hero.Name
	}

	result := strings.Join(names, ",")

	return &types.PickResponse{
		Data: result,
	}, nil
}
