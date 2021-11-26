package system

import (
	"fmt"

	"github.com/nht1206/pricetracker/config"
	"github.com/nht1206/pricetracker/db"
	"github.com/nht1206/pricetracker/internal/repository"
	"github.com/nht1206/pricetracker/internal/service/notifier"
)

type Context struct {
	Config          *config.Config
	Dao             repository.DAO
	NotifierFactory notifier.NotifierFactory
}

func InitSystemContext(cfg *config.Config) (*Context, error) {
	if cfg == nil {
		return nil, fmt.Errorf("cfg is nil")
	}

	db, err := db.InitDatabase(cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database. %v", err)
	}

	dao, err := repository.NewDAO(db)
	if err != nil {
		return nil, err
	}

	notifierFactory, err := notifier.NewNotifierFactory(cfg.Notifier)
	if err != nil {
		return nil, err
	}

	return &Context{
		Config:          cfg,
		Dao:             dao,
		NotifierFactory: notifierFactory,
	}, nil
}
