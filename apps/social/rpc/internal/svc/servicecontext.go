package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"goim/apps/social/rpc/internal/config"
	"goim/apps/social/socialmodel"
)

type ServiceContext struct {
	Config config.Config

	socialmodel.FriendsModel
	socialmodel.FriendRequestsModel
	socialmodel.GroupsModel
	socialmodel.GroupRequestsModel
	socialmodel.GroupMembersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,

		FriendsModel:        socialmodel.NewFriendsModel(sqlConn, c.Cache),
		FriendRequestsModel: socialmodel.NewFriendRequestsModel(sqlConn, c.Cache),
		GroupsModel:         socialmodel.NewGroupsModel(sqlConn, c.Cache),
		GroupRequestsModel:  socialmodel.NewGroupRequestsModel(sqlConn, c.Cache),
		GroupMembersModel:   socialmodel.NewGroupMembersModel(sqlConn, c.Cache),
	}
}
