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
	Sbjname   svc.Sbjjson
	Teacher   svc.Teachjson
	Src       svc.Srcjson
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
		var subject svc.ClassSubject
		l.svcCtx.Mysql.Where("class_code = ? AND weekly = ?", req.Cookie, int(time.Now().Weekday())).First(&subject)
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
		return &types.GetSubjectResp{
			Weekly:    time.Now().Weekday().String(),
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
	} else {
		return &types.GetSubjectResp{
			Weekly:    time.Now().Weekday().String(),
			Lession1:  sbj.Sbjname.Lession1,
			Src1:      sbj.Src.Src1,
			Teacher1:  sbj.Teacher.Teacher1,
			Lession2:  sbj.Sbjname.Lession2,
			Src2:      sbj.Src.Src2,
			Teacher2:  sbj.Teacher.Teacher2,
			Lession3:  sbj.Sbjname.Lession3,
			Src3:      sbj.Src.Src3,
			Teacher3:  sbj.Teacher.Teacher3,
			Lession4:  sbj.Sbjname.Lession4,
			Src4:      sbj.Src.Src4,
			Teacher4:  sbj.Teacher.Teacher4,
			Lession5:  sbj.Sbjname.Lession5,
			Src5:      sbj.Src.Src5,
			Teacher5:  sbj.Teacher.Teacher5,
			Lession6:  sbj.Sbjname.Lession6,
			Src6:      sbj.Src.Src6,
			Teacher6:  sbj.Teacher.Teacher6,
			Lession7:  sbj.Sbjname.Lession7,
			Src7:      sbj.Src.Src7,
			Teacher7:  sbj.Teacher.Teacher7,
			Lession8:  sbj.Sbjname.Lession8,
			Src8:      sbj.Src.Src8,
			Teacher8:  sbj.Teacher.Teacher8,
			Lession9:  sbj.Sbjname.Lession9,
			Src9:      sbj.Src.Src9,
			Teacher9:  sbj.Teacher.Teacher9,
			Lession10: sbj.Sbjname.Lession10,
			Src10:     sbj.Src.Src10,
			Teacher10: sbj.Teacher.Teacher10,
			Lession11: sbj.Sbjname.Lession11,
			Src11:     sbj.Src.Src11,
			Teacher11: sbj.Teacher.Teacher11,
			Lession12: sbj.Sbjname.Lession12,
			Src12:     sbj.Src.Src12,
			Teacher12: sbj.Teacher.Teacher12,
			Lession13: sbj.Sbjname.Lession13,
			Src13:     sbj.Src.Src13,
			Teacher13: sbj.Teacher.Teacher13,
			Lession14: sbj.Sbjname.Lession14,
			Src14:     sbj.Src.Src14,
			Teacher14: sbj.Teacher.Teacher14,
		}, nil
	}
}
