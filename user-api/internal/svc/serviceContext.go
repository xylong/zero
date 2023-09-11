package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero/user-api/internal/config"
	"zero/user-api/model"
)

type ServiceContext struct {
	Config config.Config
	model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DSN)),
	}
}
