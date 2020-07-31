package dao

import "goblog/internal/model"

// 获取认证信息
func (d *Dao) GetAuth(appKey, appSecret string) (*model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
