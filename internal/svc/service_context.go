package svc

import (
	"lottery/internal/config"
	"lottery/internal/model"
)

type ServiceContext struct {
	Config config.Config

	HeroModel model.HeroModel
	LogModel  model.LogModel
	TeamModel model.TeamModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := c.Mysql.Conn()

	return &ServiceContext{
		Config:    c,
		HeroModel: model.NewHeroModel(conn, c.Cache),
		TeamModel: model.NewTeamModel(conn),
		LogModel:  model.NewLogModel(conn),
	}
}
