package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Conf = &configuration{}

type Config struct {
	viper *viper.Viper
}

type configuration struct {
	Downloader Downloader `mapstructure:"downloader" json:"downloader" yaml:"downloader"`
	Provider   Provider   `mapstructure:"provider" json:"provider" yaml:"provider"`
	Logger        LoggerConfig        `mapstructure:"logger" json:"logger" yaml:"logger"`
}

type Downloader struct {
	Thunder Thunder `mapstructure:"thunder" json:"thunder" yaml:"thunder"`
}

type Provider struct {
	DoMP4 DoMP4 `mapstructure:"domp4" json:"domp4" yaml:"domp4"`
}

type Thunder struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type DoMP4 struct {
	Xpath          string `mapstructure:"xpath" json:"xpath" yaml:"xpath"`
	CurrentEpXpath string `mapstructure:"current_ep_xpath" json:"current_ep_xpath" yaml:"current_ep_xpath"`
}

type LoggerConfig struct {
	DebugFileName string `mapstructure:"debugFileName" json:"debugFileName" yaml:"debugFileName"`
	InfoFileName  string `mapstructure:"infoFileName" json:"infoFileName" yaml:"infoFileName"`
	WarnFileName  string `mapstructure:"warnFileName" json:"warnFileName" yaml:"warnFileName"`
	ErrorFileName string `mapstructure:"errorFileName" json:"errorFileName" yaml:"errorFileName"`
	MaxSize       int    `mapstructure:"maxSize" json:"maxSize" yaml:"maxSize"`
	MaxAge        int    `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`
	MaxBackups    int    `mapstructure:"maxBackups" json:"maxBackups" yaml:"maxBackups"`
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
