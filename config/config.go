package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"runtime"
	"time"
)

type Config struct {
	App   AppConfig
	MySQL MySQLConfig
	Redis RedisConfig
}

type AppConfig struct {
	Env              string
	Host             string
	Port             int
	Debug            bool
	URL              string
	APIBaseURL       string
	Token            string
	WorkerPoolsCount int
}

type MySQLConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Database    string
	Driver      string
	MaxOpen     int
	MaxIdle     int
	MaxLifeTime time.Duration
}

type RedisConfig struct {
	Host            string
	Port            int
	Password        string
	DefaultDatabase int
	PoolSize        int
	PoolTimeout     time.Duration
	Expiration      time.Duration
}

var (
	configInstance *Config
	absPath        string
)

func Setup() error {
	var confPath string
	cfg := os.Getenv("ENVIRONMENT")
	if len(cfg) == 0 {
		cfg = "base"
	}

	viper.SetConfigName(cfg)
	viper.SetConfigType("toml")
	if absPath == "" {
		confPath = configPath()
	} else {
		confPath = absPath
	}
	viper.AddConfigPath(confPath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}

func configPath() string {
	_, fullFilename, _, _ := runtime.Caller(0)
	pwd := path.Dir(fullFilename)
	return fmt.Sprintf("%s%s", pwd, "/")
}

func Load() *Config {
	if configInstance != nil {
		return configInstance
	}
	configInstance = &Config{
		App: AppConfig{
			Env:              viper.GetString("APP.ENV"),
			Host:             viper.GetString("APP.HOST"),
			Port:             viper.GetInt("APP.PORT"),
			Debug:            viper.GetBool("APP.DEBUG"),
			URL:              viper.GetString("APP.URL"),
			APIBaseURL:       viper.GetString("APP.API_BASE_URL"),
			Token:            viper.GetString("APP.TOKEN"),
			WorkerPoolsCount: viper.GetInt("APP.WORKER_POOLS_COUNT"),
		},
		MySQL: MySQLConfig{
			Host:        viper.GetString("MYSQL.HOST"),
			Port:        viper.GetInt("MYSQL.PORT"),
			Username:    viper.GetString("MYSQL.USERNAME"),
			Password:    viper.GetString("MYSQL.PASSWORD"),
			Database:    viper.GetString("MYSQL.DATABASE"),
			Driver:      viper.GetString("MYSQL.DRIVER"),
			MaxOpen:     viper.GetInt("MYSQL.MAX_OPEN_CONNECTIONS"),
			MaxIdle:     viper.GetInt("MYSQL.MAX_IDLE_CONNECTIONS"),
			MaxLifeTime: viper.GetDuration("MYSQL.CONN_MAX_LIFE_TIME") * time.Minute,
		},
		Redis: RedisConfig{
			Host:            viper.GetString("REDIS.HOST"),
			Port:            viper.GetInt("REDIS.PORT"),
			Password:        viper.GetString("REDIS.PASSWORD"),
			DefaultDatabase: viper.GetInt("REDIS.DATABASE"),
			PoolSize:        viper.GetInt("REDIS.POOL_SIZE"),
			PoolTimeout:     viper.GetDuration("REDIS.POOL_TIMEOUT"),
			Expiration:      viper.GetDuration("REDIS.EXPIRATION"),
		},
	}
	return configInstance
}
