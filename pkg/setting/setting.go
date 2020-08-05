package setting

import "github.com/spf13/viper"

// setting 程序配置组件
type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	// 设置配置文件名，配置文件路径，配置文件格式
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
