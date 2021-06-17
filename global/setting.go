package global

import (
	"selfblog/pkg/logger"
	"selfblog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
	EmailSetting    *setting.EmailSettingS
)

