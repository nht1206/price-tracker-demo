package pricetracker

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nht1206/pricetracker/config"
	"github.com/nht1206/pricetracker/internal/log"
	"github.com/nht1206/pricetracker/internal/model"
	"github.com/nht1206/pricetracker/internal/repository"
	"github.com/nht1206/pricetracker/internal/service/notifier"
	"github.com/nht1206/pricetracker/system"
)

const (
	TestAppName                = "dummy_app_name"
	TestNumCrawlingGoroutines  = 1
	TestNumNotifyingGoroutines = 1
	TestLogLevel               = "info"
	TestDSN                    = "dummy_dsn"
	TestMailSMTPHost           = "dummy_smtp_host"
	TestMailSMTPPort           = "dummy_smtp_port"
	TestMailSender             = "dummy_sender"
	TestMailSenderPassword     = "dummy_sender_password"
)

func TestStartApp(t *testing.T) {

	cfg := &config.Config{
		AppName:                TestAppName,
		NumCrawlingGoroutines:  TestNumCrawlingGoroutines,
		NumNotifyingGoroutines: TestNumNotifyingGoroutines,
		Log: &config.LogConfig{
			Level: TestLogLevel,
		},
		DB: &config.DatabaseConfig{
			DSN: TestDSN,
		},
		Notifier: &config.NotifierConfig{
			Mail: &config.MailConfig{
				SMTPHost:       TestMailSMTPHost,
				SMTPPort:       TestMailSMTPPort,
				Sender:         TestMailSender,
				SenderPassword: TestMailSenderPassword,
			},
		},
	}

	logger, err := log.InitLogger(cfg.Log)
	if err != nil {
		t.Error(err)
	}
	params := []struct {
		newDao             func() *repository.MockDAO
		newNotifierFactory func() *notifier.MockNotifierFactory
	}{
		{
			newDao: func() *repository.MockDAO {
				ctrl := gomock.NewController(t)
				m := repository.NewMockDAO(ctrl)
				m.EXPECT().FindTargetTrackingProduct().Return([]model.Product{}, nil)
				return m
			},
			newNotifierFactory: func() *notifier.MockNotifierFactory {
				ctrl := gomock.NewController(t)
				fnm := notifier.NewMockNotifierFactory(ctrl)
				return fnm
			},
		},
	}

	for _, p := range params {
		dao := p.newDao()
		notifierFactory := p.newNotifierFactory()

		sysCtx := &system.Context{
			Config:          cfg,
			Dao:             dao,
			NotifierFactory: notifierFactory,
		}
		ctx := log.WithLogger(context.Background(), logger)

		StartApp(ctx, sysCtx)
	}
}
