package logic

import (
	"context"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeSubjLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeSubjLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeSubjLogic {
	return &ChangeSubjLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeSubjLogic) ChangeSubj(req *types.ChangeSubjReq) (resp *types.ChangeSubjResp, err error) {
	l.svcCtx.Redis.Del(req.ClassCode)
	//var sbj svc.Subject
	subject := &Subject{
		ClassCode: req.ClassCode,
		Weekly:    req.Weekly,
		Sbjname: svc.Sbjjson{
			Lession1:  req.Lession1,
			Lession2:  req.Lession2,
			Lession3:  req.Lession3,
			Lession4:  req.Lession4,
			Lession5:  req.Lession5,
			Lession6:  req.Lession6,
			Lession7:  req.Lession7,
			Lession8:  req.Lession8,
			Lession9:  req.Lession9,
			Lession10: req.Lession10,
			Lession11: req.Lession11,
			Lession12: req.Lession12,
			Lession13: req.Lession13,
		},
		Teacher: svc.Teachjson{
			Teacher1:  req.Teacher1,
			Teacher2:  req.Teacher2,
			Teacher3:  req.Teacher3,
			Teacher4:  req.Teacher4,
			Teacher5:  req.Teacher5,
			Teacher6:  req.Teacher6,
			Teacher7:  req.Teacher7,
			Teacher8:  req.Teacher8,
			Teacher9:  req.Teacher9,
			Teacher10: req.Teacher10,
			Teacher11: req.Teacher11,
			Teacher12: req.Teacher12,
			Teacher13: req.Teacher13,
		},
		Src: svc.Srcjson{
			Src1:  "",
			Src2:  "",
			Src3:  "",
			Src4:  "",
			Src5:  "",
			Src6:  "",
			Src7:  "",
			Src8:  "",
			Src9:  "",
			Src10: "",
			Src11: "",
			Src12: "",
			Src13: "",
		},
	}
	l.svcCtx.Mysql.Table("subjects").Where("class_code = ? AND weekly = ?", req.ClassCode, req.Weekly).Updates(subject)
	return &types.ChangeSubjResp{
		Status: "ok",
	}, nil
}