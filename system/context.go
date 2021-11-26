package system

import (
	"fmt"

	"github.com/nht1206/pricetracker/config"
	"github.com/nht1206/pricetracker/db"
	"github.com/nht1206/pricetracker/internal/repository"
)

type Context struct {
	Config      *config.Config
	ProductRepo repository.ProductRepository
}

func InitSystemContext(cfg *config.Config) (*Context, error) {
	if cfg == nil {
		return nil, fmt.Errorf("cfg is nil")
	}

	db, err := db.InitDatabase(cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database. %v", err)
	}

	productRepo, err := repository.NewProductRepository(db)
	if err != nil {
		return nil, err
	}
	return &Context{
		Config:      cfg,
		ProductRepo: productRepo,
	}, nil
}
