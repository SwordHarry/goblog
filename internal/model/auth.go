package model

import (
	"github.com/jinzhu/gorm"
	"goblog/internal/dao"
)

// 获取认证信息
// 根据传入的 app_key 和 app_secret 进行验证，判断是否存在这样一条数据
func (m *Model) GetAuth(appKey, appSecret string) (*dao.Auth, error) {
	auth := dao.Auth{AppKey: appKey, AppSecret: appSecret}
	db := m.engine
	db = db.Where("app_key = ? and app_secret = ? and is_del = ?", appKey, appSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &auth, err
}
