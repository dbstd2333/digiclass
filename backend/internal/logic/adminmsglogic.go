package logic

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminMsgLogic {
	return &AdminMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminMsgLogic) AdminMsg(req *types.AdminMsgReq) (resp *types.AdminMsgResp, err error) {
	err = l.svcCtx.Redis.Del("classmsg").Err()
	if err != nil {
		logs := svc.Log{
			Type: "warning",
			Logs: "redis del key err: " + err.Error(),
			Time: time.Now(),
		}
		l.svcCtx.Mysql.Create(&logs)
	}
	var msg svc.Message
	var msg2 svc.Message
	msg = svc.Message{
		ClassCode: req.ClassCode,
		Msg:       req.Msg,
		MsgTitle:  req.MsgTitle,
		Time:      time.Now().Local(),
	}
	l.svcCtx.Mysql.Model(msg2).Table("messages").Create(&msg)
	l.svcCtx.Mysql.Model(msg2).Table("messages").Last(&msg2)
	jmsg, err := json.Marshal(&msg2)
	err2 := l.svcCtx.Redis.Set("classmsg", jmsg, 28800*time.Second).Err()
	if err != nil || err2 != nil {
		return &types.AdminMsgResp{
			Status: "err:" + err.Error() + "err redis:" + err2.Error(),
		}, nil
	} else {
		return &types.AdminMsgResp{
			Status: "ok",
		}, nil
	}
}
