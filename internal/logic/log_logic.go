package logic

import (
	"context"
	"errors"
	"lottery/internal/model"
	"lottery/internal/svc"
	"lottery/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Pick a team by encryptCode
func NewLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogLogic {
	return &LogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogLogic) Log(req *types.LogRequest) (resp *types.LogResponse, err error) {
	_, err = l.svcCtx.TeamModel.FindOneByEncryptCode(l.ctx, req.EncryptCode)
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

	return &types.LogResponse{}, nil
}
