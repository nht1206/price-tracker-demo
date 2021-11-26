package notifier

import (
	"fmt"

	"github.com/nht1206/pricetracker/config"
	"github.com/nht1206/pricetracker/internal/model"
)

type NotifierType string

const (
	Mail NotifierType = "mail"
)

type NotifierFactory interface {
	CreateNotifier(followType uint) (Notifier, error)
}

type notifierFactory struct {
	config *config.NotifierConfig
}

func NewNotifierFactory(cfg *config.NotifierConfig) (NotifierFactory, error) {
	if cfg == nil {
		return nil, fmt.Errorf("cfg is nil")
	}

	return &notifierFactory{
		config: cfg,
	}, nil
}

func (f *notifierFactory) CreateNotifier(followType uint) (Notifier, error) {
	switch followType {
	case 1:
		return newMailNotifier(f.config.Mail)
	default:
		return nil, fmt.Errorf("notifier not found")
	}
}

type Notifier interface {
	Notify(user *model.User, result *model.TrackingResult) error
}
