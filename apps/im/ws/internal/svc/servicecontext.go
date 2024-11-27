package svc

import (
	immodels "goim/apps/im/immodel"
	"goim/apps/im/ws/internal/config"
	"goim/apps/task/mqclient"
)

type ServiceContext struct {
	Config config.Config

	immodels.ChatLogModel
	mqclient.MsgChatTransferClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		MsgChatTransferClient: mqclient.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
		ChatLogModel:          immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}
}
