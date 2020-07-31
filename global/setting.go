package global

import (
	"goblog/pkg/logger"
	"goblog/pkg/setting"
)

// 全局配置对象
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
)

const DebugMode = "debug"
