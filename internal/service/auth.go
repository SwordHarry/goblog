package service

import (
	"errors"
	"goblog/internal/request"
)

func (svc *Service) CheckAuth(param *request.AuthRequest) error {
	auth, err := svc.model.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
