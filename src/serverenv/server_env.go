package serverenv

import (
	log "github.com/sirupsen/logrus"
	"github.com/trangnkp/my_books/src/config"
	"github.com/trangnkp/my_books/src/service"
	"github.com/trangnkp/my_books/src/store"
)

type ServerEnv struct {
	Config              *config.AppConfig
	DBStores            *store.DBStores
	NotificationService *service.KafkaNotificationService
}

func NewServerEnv(cfg *config.AppConfig) (*ServerEnv, error) {
	stores, err := store.NewDBStores(cfg.DB)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	notificationSvc, err := service.NewKafkaNotificationService(cfg.KafkaProducerConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &ServerEnv{
		Config:              cfg,
		DBStores:            stores,
		NotificationService: notificationSvc,
	}, nil
}
