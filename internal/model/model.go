package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goblog/global"
	"goblog/pkg/setting"
)

// 数据库链接组件
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == global.DebugMode {
		// 打印 log 信息
		db.LogMode(true)
	}
	// 默认使用单表
	db.SingularTable(true)
	dbp := db.DB()
	// 设置连接池中的连接数
	dbp.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	// 设置打开的最大连接数
	dbp.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}
