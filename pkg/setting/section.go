package setting

import "time"

// ServerSettings 服务配置
type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

// AppSettings APP配置
type AppSettings struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

// DatabaseSettings 数据库配置
type DatabaseSettings struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePreFix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// ReadSection 读取配置
func (s *Setting) ReadSection(k string, v any) (err error) {
	err = s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return
}
