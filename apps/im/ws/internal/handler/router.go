package handler

import (
	"goim/apps/im/ws/internal/handler/conversation"
	"goim/apps/im/ws/internal/handler/push"
	"goim/apps/im/ws/internal/handler/user"
	"goim/apps/im/ws/internal/svc"
	"goim/apps/im/ws/websocket"
)

func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoutes([]websocket.Route{
		{
			Method:  "user.online",
			Handler: user.OnLine(svc),
		},
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "push",
			Handler: push.Push(svc),
		},
	})
}
