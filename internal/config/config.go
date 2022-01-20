package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var C config

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
	err2 := viper.Unmarshal(&C)
	if err != nil {
		return err2
	}
	watchConfig()

	return nil
}

func watchConfig() {
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.Unmarshal(&C)
	})
	viper.WatchConfig()
}

type config struct {
	Port uint32
	Domain string
	Mysql mysql
}
type mysql struct {
	Addr string
	username string
	password string
	db string
	max_idle_conn uint32
	max_open_conn uint32
	conn_max_life_time uint32
}