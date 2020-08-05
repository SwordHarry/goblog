package service

import (
	"context"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"goblog/global"
	"goblog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
	}
	// 数据库链接信息上下文注册
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}
