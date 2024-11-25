package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"goim/apps/user/model"
	"goim/apps/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUsersModel(sqlConn, c.Cache),
	}
}
