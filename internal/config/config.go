package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf Config

func Init(configPath string) error {
	err := initConfig(configPath)
	if err != nil {
		return err
	}

	return nil
}

func initConfig(path string) error {
	if path == "" {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config.local")
	} else {
		viper.SetConfigFile(path)
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err2 := viper.Unmarshal(&Conf)
	if err != nil {
		return err2
	}
	watchConfig()

	return nil
}

func watchConfig() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.Unmarshal(&Conf)
	})
	viper.WatchConfig()
}

type Config struct {
	Port uint32
	Domain string
	Mysql mysql
}
type mysql struct {
	Addr string
	Username string
	Password string
	DB string
	MaxIdleConns int
	MaxOpenConns int
	ConnMaxLifeTime int
}