package system

import (
	"testing"

	"github.com/nht1206/pricetracker/config"
)

const (
	TestAppName                = "dummy_app_name"
	TestNumCrawlingGoroutines  = 1
	TestNumNotifyingGoroutines = 1
	TestLogLevel               = "info"
	TestDSN                    = "root:@tcp(127.0.0.1:3306)/test_db"
	TestMailSMTPHost           = "dummy_smtp_host"
	TestMailSMTPPort           = "dummy_smtp_port"
	TestMailSender             = "dummy_sender"
	TestMailSenderPassword     = "dummy_sender_password"
)

func TestInitSystemContext(t *testing.T) {
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

	_, err := InitSystemContext(cfg)

	if err != nil {
		t.Error("err is not nil")
	}
}

func TestNGInitSystemContext(t *testing.T) {
	t.Run("Case 1: Failed to init system context. Because config is nil", func(t *testing.T) {
		_, err := InitSystemContext(nil)

		if err == nil {
			t.Error("err is nil")
		}
	})

	t.Run("Case 1: Failed to init system context. Because db.DB is not initialized", func(t *testing.T) {
		_, err := InitSystemContext(&config.Config{
			AppName:                "Test",
			NumCrawlingGoroutines:  10,
			NumNotifyingGoroutines: 5,
			Log: &config.LogConfig{
				OutputPath: "test",
				FileName:   "test.log",
				Level:      "info",
			},
			DB: &config.DatabaseConfig{
				DSN: "",
			},
		})

		if err == nil {
			t.Error("err is nil")
		}
	})
}
