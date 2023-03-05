package logic

import (
	"context"
	"encoding/json"
	"time"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type Message struct {
	Uuid      int64
	ClassCode string
	Msg       string
	MsgTitle  string
	Sender    string
	Time      time.Time
}

func NewClassMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassMsgLogic {
	return &ClassMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassMsgLogic) ClassMsg(req *types.ClassMsgReq) (resp *types.ClassMsgResp, err error) {
	classmsg, err := l.svcCtx.Redis.Get("classmsg").Result()
	cmsg := Message{}
	errj := json.Unmarshal([]byte(classmsg), &cmsg)
	if errj != nil {
		logs := svc.Log{
			Type: "warning",
			Logs: "class get msg json / get redis cache err: " + errj.Error(),
			Time: time.Now(),
		}
		l.svcCtx.Mysql.Create(&logs)
		var msg svc.Message
		l.svcCtx.Mysql.Model(msg).Table("messages").Last(&msg)
		return &types.ClassMsgResp{
			Msgid:     msg.Uuid,
			Msg:       msg.Msg,
			MsgTitle:  msg.MsgTitle,
			ClassCode: msg.ClassCode,
			Time:      msg.Time.String(),
			Sender:    msg.Sender,
		}, nil
	} else {
		return &types.ClassMsgResp{
			Msgid:     cmsg.Uuid,
			Msg:       cmsg.Msg,
			MsgTitle:  cmsg.MsgTitle,
			ClassCode: cmsg.ClassCode,
			Time:      cmsg.Time.String(),
			Sender:    cmsg.Sender,
		}, nil
	}
}
