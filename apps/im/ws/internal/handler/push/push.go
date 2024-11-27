package push

import (
	"github.com/mitchellh/mapstructure"
	"goim/apps/im/ws/internal/svc"
	"goim/apps/im/ws/websocket"
	"goim/apps/im/ws/wsmodel"
)

func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data wsmodel.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err))
			return
		}

		// 发送的目标
		rconn := srv.GetConn(data.RecvId)
		if rconn == nil {
			// todo: 目标离线
			return
		}

		srv.Infof("push msg %v", data)

		srv.Send(websocket.NewMessage(data.SendId, &wsmodel.Chat{
			ConversationId: data.ConversationId,
			ChatType:       data.ChatType,
			SendTime:       data.SendTime,
			Msg: wsmodel.Msg{
				MType:   data.MType,
				Content: data.Content,
			},
		}), rconn)
	}
}
