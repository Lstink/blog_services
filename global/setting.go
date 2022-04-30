package global

import (
	"github.com/lstink/blog/pkg/logger"
	"github.com/lstink/blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)
