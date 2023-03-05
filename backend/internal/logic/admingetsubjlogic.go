package logic

import (
	"context"
	"strconv"
	"time"

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
	week, err := strconv.Atoi(req.Weekly)
	err = l.svcCtx.Redis.Del(req.ClassCode + time.Weekday(week).String()).Err()
	if err != nil {
		logs := svc.Log{
			Type: "warning",
			Logs: "redis del key err: " + err.Error(),
			Time: time.Now(),
		}
		l.svcCtx.Mysql.Create(&logs)
	}
	var subject svc.ClassSubject
	l.svcCtx.Mysql.Where("class_code = ? AND weekly = ?", req.ClassCode, week).First(&subject)
	return &types.AdminGetSubjResp{
		Lession1:  subject.Sbjname.Lession1,
		Src1:      subject.Src.Src1,
		Teacher1:  subject.Teacher.Teacher1,
		Lession2:  subject.Sbjname.Lession2,
		Src2:      subject.Src.Src2,
		Teacher2:  subject.Teacher.Teacher2,
		Lession3:  subject.Sbjname.Lession3,
		Src3:      subject.Src.Src3,
		Teacher3:  subject.Teacher.Teacher3,
		Lession4:  subject.Sbjname.Lession4,
		Src4:      subject.Src.Src4,
		Teacher4:  subject.Teacher.Teacher4,
		Lession5:  subject.Sbjname.Lession5,
		Src5:      subject.Src.Src5,
		Teacher5:  subject.Teacher.Teacher5,
		Lession6:  subject.Sbjname.Lession6,
		Src6:      subject.Src.Src6,
		Teacher6:  subject.Teacher.Teacher6,
		Lession7:  subject.Sbjname.Lession7,
		Src7:      subject.Src.Src7,
		Teacher7:  subject.Teacher.Teacher7,
		Lession8:  subject.Sbjname.Lession8,
		Src8:      subject.Src.Src8,
		Teacher8:  subject.Teacher.Teacher8,
		Lession9:  subject.Sbjname.Lession9,
		Src9:      subject.Src.Src9,
		Teacher9:  subject.Teacher.Teacher9,
		Lession10: subject.Sbjname.Lession10,
		Src10:     subject.Src.Src10,
		Teacher10: subject.Teacher.Teacher10,
		Lession11: subject.Sbjname.Lession11,
		Src11:     subject.Src.Src11,
		Teacher11: subject.Teacher.Teacher11,
		Lession12: subject.Sbjname.Lession12,
		Src12:     subject.Src.Src12,
		Teacher12: subject.Teacher.Teacher12,
		Lession13: subject.Sbjname.Lession13,
		Src13:     subject.Src.Src13,
		Teacher13: subject.Teacher.Teacher13,
		Lession14: subject.Sbjname.Lession14,
		Src14:     subject.Src.Src14,
		Teacher14: subject.Teacher.Teacher14,
	}, nil
}
