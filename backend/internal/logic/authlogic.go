package logic

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Auth(req *types.AuthReq) (resp *types.AuthResp, err error) {
	var auth svc.Userinf
	err = l.svcCtx.Mysql.Table("userinfs").Where("class_code = ? AND passwd = ?", req.AuthCode, req.Passwd).First(&auth).Error
	if err != nil {
		return &types.AuthResp{
			Status: "error",
			Cookie: "error",
		}, nil
	} else {
		return &types.AuthResp{
			Status: "ok",
			Cookie: auth.ClassCode,
		}, nil
	}
}
