package container

import (
	"fmt"
	"github.com/alihdrii/go-hexagonal/internal/adapter/http/handler"
	"github.com/alihdrii/go-hexagonal/pkg/setup/mysql"
	"github.com/alihdrii/go-hexagonal/pkg/setup/redis"
	"log"
)

type AppContainer struct {
	Core     *CoreContainer
	DB       *mysql.DB
	Redis    *redis.Client
	Handlers *AppContainerHandler
}

type AppContainerHandler struct {
	User *handler.UserHandler
	// Add more as needed
}

func NewAppContainer() (*AppContainer, error) {
	core, err := NewCoreContainer()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Core config: %w", err)
	}

	mysqlDB, err := setupMysql(core.Config.MySQL)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize MySQL: %w", err)
	}

	rDB, err := setupRedis(core.Config.Redis)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Redis: %w", err)
	}

	return &AppContainer{
		Core:     core,
		DB:       mysqlDB,
		Redis:    rDB,
		Handlers: &AppContainerHandler{},
	}, nil
}

func (c *AppContainer) CloseContainer() {
	if err := c.DB.Close(); err != nil {
		log.Printf("Failed to close MySQL connection: %v", err)
	}

	if err := c.Redis.Close(); err != nil {
		log.Printf("Failed to close Redis connection: %v", err)
	}
}
