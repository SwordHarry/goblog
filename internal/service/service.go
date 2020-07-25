package service

import (
	"context"
	"goblog/global"
	"goblog/internal/dao"
)

type Service struct {
	ctx *context.Context
	dao *dao.Dao
}

func New(ctx *context.Context) Service {
	return Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
}
