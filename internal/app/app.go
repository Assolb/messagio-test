package app

import (
	"github.com/sirupsen/logrus"
	"messagio/internal/config"
	"messagio/internal/repository"
	"messagio/internal/service"
	"messagio/internal/transport"
)

func Run() {
	err := config.LoadConfig("configs/config.yml")
	if err != nil {
		logrus.Fatalf("Error loading config: %v", err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		logrus.Fatalf("Error loading config: %v", err)
	}

	db, err := repository.NewDatabase(&cfg.Database)
	if err != nil {
		logrus.Fatalf("Error initializing database: %v", err)
	}

	repo, err := repository.NewRepository(db)
	if err != nil {
		logrus.Fatalf("Error initializing repository: %v", err)
	}

	svc := service.NewService(repo, cfg)

	err = transport.InitRouter(svc)
	if err != nil {
		logrus.Fatalf("Error initializing controller: %v", err)
		return
	}

	defer db.Db.Close()
}
