package logic

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetSubjLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetSubjLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetSubjLogic {
	return &AdminGetSubjLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetSubjLogic) AdminGetSubj(req *types.AdminGetSubjReq) (resp *types.AdminGetSubjResp, err error) {
	// todo: add your logic here and delete this line

	return
}
