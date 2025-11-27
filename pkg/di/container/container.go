package container

import (
	"fmt"
	"github.com/alihdrii/go-hexagonal/config"
	"github.com/alihdrii/go-hexagonal/pkg/setup/mysql"

	"time"

	rdb "github.com/alihdrii/go-hexagonal/pkg/setup/redis"
)

type MainContainerI interface {
	CloseContainer()
}

func setupMysql(cfg config.MySQLConfig) (*mysql.DB, error) {
	mysqlCfg := mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.Username,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database),
		MaxOpen:     cfg.MaxOpen,
		MaxIdle:     cfg.MaxIdle,
		MaxLifetime: time.Minute * cfg.MaxLifeTime,
	}

	dbConn, err := mysql.NewConnection(mysqlCfg)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func setupRedis(cfg config.RedisConfig) (*rdb.Client, error) {
	redisAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	redisCfg := rdb.Config{
		Addr:        redisAddr,
		Password:    cfg.Password,
		Database:    cfg.DefaultDatabase,
		PoolSize:    cfg.PoolSize,
		PoolTimeout: cfg.PoolTimeout,
	}

	client, err := rdb.NewRedis(redisCfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}
