package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"goim/apps/user/rpc/internal/config"
	"goim/apps/user/usermodel"
)

type ServiceContext struct {
	Config config.Config
	usermodel.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: usermodel.NewUsersModel(sqlConn, c.Cache),
	}
}
