package config

import (
	"log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = &configuration{}

type Config struct {
	viper *viper.Viper
}

type configuration struct {
	Downloader Downloader `mapstructure:"downloader" json:"downloader" yaml:"downloader"`
}

type Downloader struct {
	Thunder Thunder `mapstructure:"thunder" json:"thunder" yaml:"thunder"`
}

type Thunder struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func InitConfig() *Config {
	config := &Config{
		viper: viper.New(),
	}
	config.viper.SetConfigName("config")
	config.viper.AddConfigPath("./config")
	config.viper.SetConfigType("yaml")

	if err := config.viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %v\n", err)
	}
	config.viper.WatchConfig()
	config.viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file changed:", in.Name)
		if err := config.viper.Unmarshal(Conf); err != nil {
			log.Println("Unmarshal config failed, err:", err)
		}
	})
	if err := config.viper.Unmarshal(Conf); err != nil {
		log.Fatalf("Unmarshal config failed, err:%v\n", err)
	}
	return config
}
