package logic

import (
	"context"
	"encoding/json"
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
	Weekday, err := strconv.Atoi(req.Weekly)
	err = l.svcCtx.Redis.Del(req.ClassCode + time.Weekday(Weekday).String()).Err()
	if err != nil {
		logs := svc.Log{
			Type: "warning",
			Logs: "redis del key err: " + err.Error(),
			Time: time.Now(),
		}
		l.svcCtx.Mysql.Create(&logs)
	}
	var subject svc.ClassSubject2
	l.svcCtx.Mysql.Table("class_subjects").Where("class_code = ? AND weekly = ?", req.ClassCode, Weekday).First(&subject)
	var sbjjson svc.Sbjjson
	err = json.Unmarshal([]byte(subject.Sbjname), &sbjjson)
	var teacher svc.Teachjson
	err = json.Unmarshal([]byte(subject.Teacher), &teacher)
	var src svc.Srcjson
	err = json.Unmarshal([]byte(subject.Src), &src)
	return &types.AdminGetSubjResp{
		Lession1:  sbjjson.Lession1,
		Src1:      src.Src1,
		Teacher1:  teacher.Teacher1,
		Lession2:  sbjjson.Lession2,
		Src2:      src.Src2,
		Teacher2:  teacher.Teacher2,
		Lession3:  sbjjson.Lession3,
		Src3:      src.Src3,
		Teacher3:  teacher.Teacher3,
		Lession4:  sbjjson.Lession4,
		Src4:      src.Src4,
		Teacher4:  teacher.Teacher4,
		Lession5:  sbjjson.Lession5,
		Src5:      src.Src5,
		Teacher5:  teacher.Teacher5,
		Lession6:  sbjjson.Lession6,
		Src6:      src.Src6,
		Teacher6:  teacher.Teacher6,
		Lession7:  sbjjson.Lession7,
		Src7:      src.Src7,
		Teacher7:  teacher.Teacher7,
		Lession8:  sbjjson.Lession8,
		Src8:      src.Src8,
		Teacher8:  teacher.Teacher8,
		Lession9:  sbjjson.Lession9,
		Src9:      src.Src9,
		Teacher9:  teacher.Teacher9,
		Lession10: sbjjson.Lession10,
		Src10:     src.Src10,
		Teacher10: teacher.Teacher10,
		Lession11: sbjjson.Lession11,
		Src11:     src.Src11,
		Teacher11: teacher.Teacher11,
		Lession12: sbjjson.Lession12,
		Src12:     src.Src12,
		Teacher12: teacher.Teacher12,
		Lession13: sbjjson.Lession13,
		Src13:     src.Src13,
		Teacher13: teacher.Teacher13,
		Lession14: sbjjson.Lession14,
		Src14:     src.Src14,
		Teacher14: teacher.Teacher14,
	}, nil
}
