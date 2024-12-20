package logic

import (
	"context"
	"errors"
	"lottery/internal/model"
	"strings"
	"time"

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
		if errors.Is(err, model.ErrNotFound) {
			logx.Errorw("team not found",
				logx.Field("encrypt_code", req.EncryptCode),
				logx.Field("error", err),
			)
			return nil, errors.New("team not found, please check your encrypt code")
		} else {
			return nil, errors.New("internal server error")
		}
	}

	if team == nil {
		logx.Error("team not found, please check your encryptCode")
		return nil, errors.New("team not found, please check your encryptCode")
	}

	if !team.PickContent.Valid {
		return &types.PickResponse{
			Data: team.PickContent.String,
		}, nil
	}

	heroes, err := l.svcCtx.HeroModel.FindGroupIsNotPick(l.ctx)
	if err != nil {
		return nil, errors.New("heroes not found")
	}

	names := make([]string, len(heroes))
	for i, hero := range heroes {
		names[i] = hero.Name
	}

	for _, hero := range heroes {
		hero.IsPick = 1
		err := l.svcCtx.HeroModel.Update(l.ctx, hero)
		if err != nil {
			return nil, errors.New("heroes update failed")
		}
	}

	result := strings.Join(names, ",")

	log := &model.Log{
		TeamId:    team.Id,
		PickGroup: result,
		Time:      time.Now(),
	}
	_, err = l.svcCtx.LogModel.Insert(l.ctx, log)
	if err != nil {
		return nil, err
	}

	return &types.PickResponse{
		Data: result,
	}, nil
}
