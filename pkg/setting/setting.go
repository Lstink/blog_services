package setting

import "github.com/spf13/viper"

// Setting 声明全局配置的结构体
type Setting struct {
	vp *viper.Viper
}

// NewSetting 配置实例
func NewSetting() (*Setting, error) {
	// 获取实例
	vp := viper.New()
	// 设置配置名称
	vp.SetConfigName("config")
	// 添加配置目录
	vp.AddConfigPath("configs/")
	// 设置文件类型
	vp.SetConfigType("yaml")
	// 读取文件
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	// 返回配置实例
	return &Setting{vp}, nil
}
