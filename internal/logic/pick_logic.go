package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
