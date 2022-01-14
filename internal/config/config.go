package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
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
	// viper.AutomaticEnv()
	env := os.Getenv("GO_ENV")
	if env != "" {
		viper.SetEnvPrefix(env)
	} else {
		viper.SetEnvPrefix("dev")
	}

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
		fmt.Println("Config file changed:", e.Name)
		viper.Unmarshal(&C)
	})
	viper.WatchConfig()
}

type config struct {
	Port int
	Domain string
}