package config

import (
	"log"
	"nasspider/pkg/constants"
	"os"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = &configuration{}

type Config struct {
	viper *viper.Viper
}

type configuration struct {
	Server     ServerConfig `mapstructure:"server" json:"server" yaml:"server"`
	Cron       CronConfig   `mapstructure:"cron" json:"cron" yaml:"cron"`
	Downloader Downloader   `mapstructure:"downloader" json:"downloader" yaml:"downloader"`
	Logger     LoggerConfig `mapstructure:"logger" json:"logger" yaml:"logger"`
	DB         MySqlConfig  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Jwt        JwtConfig    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Passport   Passport     `mapstructure:"passport" json:"passport" yaml:"passport"`
}

type ServerConfig struct {
	Port  int  `mapstructure:"port" json:"port" yaml:"port"`
	Debug bool `mapstructure:"debug" json:"debug" yaml:"debug"`
}

type CronConfig struct {
	TvTask Cron `mapstructure:"tv_task" json:"tv_task" yaml:"tv_task"`
}

type Cron struct {
	Enabled bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Spec    string `mapstructure:"spec" json:"spec" yaml:"spec"`
}

type JwtConfig struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"` // token 有效期（秒）
}

type Downloader struct {
	Thunder     Thunder     `mapstructure:"thunder" json:"thunder" yaml:"thunder"`
	QBittorrent QBittorrent `mapstructure:"qbittorrent" json:"qbittorrent" yaml:"qbittorrent"`
}

type Thunder struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

type QBittorrent struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
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

type MySqlConfig struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	Username            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}

type Passport struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
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

func GetConf[T any](def T, envKey constants.ENVConfig) T {
	if envKey == "" {
		return def
	}
	val := os.Getenv(string(envKey))
	if val == "" {
		return def
	}

	// 根据默认值类型进行相应的类型转换
	switch any(def).(type) {
	case string:
		return any(val).(T)
	case int:
		if v, err := strconv.Atoi(val); err == nil {
			return any(v).(T)
		}
	case int64:
		if v, err := strconv.ParseInt(val, 10, 64); err == nil {
			return any(v).(T)
		}
	case float64:
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			return any(v).(T)
		}
	case bool:
		if v, err := strconv.ParseBool(val); err == nil {
			return any(v).(T)
		}
	}
	return def
}
