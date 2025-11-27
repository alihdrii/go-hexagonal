package container

import (
	"github.com/alihdrii/go-hexagonal/config"
	"log"
)

type CoreContainer struct {
	Config *config.Config
}

func NewCoreContainer() (*CoreContainer, error) {
	if err := config.Setup(); err != nil {
		log.Fatalf("[CONFIG] error setting up config: %v", err)
	}

	cfg := config.Load()

	return &CoreContainer{
		Config: cfg,
	}, nil
}
