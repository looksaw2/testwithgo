package config

import (
	"flag"
	"github.com/spf13/viper"
)

// 读取viper的文件
func NewViperConfig() error {
	//利用参数读取对应的global文件路径
	configPath := flag.String("config-path", "./config", "配置文件所在目录")
	viper.SetConfigName("global")
	viper.SetConfigType("yml")
	viper.AddConfigPath(*configPath)
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}
