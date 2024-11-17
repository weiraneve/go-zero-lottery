package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
