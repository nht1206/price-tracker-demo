package system

import (
	"fmt"

	"github.com/nht1206/pricetracker/config"
	"github.com/nht1206/pricetracker/db"
	"github.com/nht1206/pricetracker/internal/repository"
)

type Context struct {
	Config *config.Config
	Dao    repository.DAO
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
	return &Context{
		Config: cfg,
		Dao:    dao,
	}, nil
}
