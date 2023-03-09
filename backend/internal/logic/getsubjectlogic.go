package logic

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type Subject struct {
	ClassCode string
	Weekly    int
	Sbjname   string
	Teacher   string
	Src       string
}

func NewGetSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubjectLogic {
	return &GetSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubjectLogic) GetSubject(req *types.GetSubjectReq) (resp *types.GetSubjectResp, err error) {
	rsbj, err := l.svcCtx.Redis.Get(req.Cookie + time.Now().Weekday().String()).Result()
	sbj := Subject{}
	errj := json.Unmarshal([]byte(rsbj), &sbj)

	if err != nil || errj != nil { //查mysql ，write redis , re data
		logs := svc.Log{
			Type: "info",
			Logs: "redis缓存未命中/json解析失败:" + err.Error() + errj.Error(),
			Time: time.Now(),
		}
		l.svcCtx.Mysql.Create(&logs)
		var subject svc.ClassSubject2
		l.svcCtx.Mysql.Table("class_subjects").Where("class_code = ? AND weekly = ?", req.Cookie, int(time.Now().Weekday())).First(&subject)
		rsubject, _ := json.Marshal(&subject)
		rwerr := l.svcCtx.Redis.Set(req.Cookie+time.Now().Weekday().String(), rsubject, 28800*time.Second).Err()
		if rwerr != nil {
			logs := svc.Log{
				Type: "warning",
				Logs: "redis緩存寫入失敗: " + rwerr.Error(),
				Time: time.Now(),
			}
			l.svcCtx.Mysql.Create(&logs)
		}
		var sbjjson svc.Sbjjson
		errj = json.Unmarshal([]byte(subject.Sbjname), &sbjjson)
		var teacher svc.Teachjson
		errj = json.Unmarshal([]byte(subject.Teacher), &teacher)
		var src svc.Srcjson
		errj = json.Unmarshal([]byte(subject.Src), &src)
		return &types.GetSubjectResp{
			Weekly:    time.Now().Weekday().String(),
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
	} else {
		var sbjjson svc.Sbjjson
		errj = json.Unmarshal([]byte(sbj.Sbjname), &sbjjson)
		var teacher svc.Teachjson
		errj = json.Unmarshal([]byte(sbj.Teacher), &teacher)
		var src svc.Srcjson
		errj = json.Unmarshal([]byte(sbj.Src), &src)
		return &types.GetSubjectResp{
			Weekly:    time.Now().Weekday().String(),
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
}
