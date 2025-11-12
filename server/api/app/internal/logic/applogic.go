// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"server/api/app/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppLogic {
	return &AppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
