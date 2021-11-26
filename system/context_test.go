package system

import (
	"testing"

	"github.com/nht1206/pricetracker/config"
)

func TestInitSystemContext(t *testing.T) {
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
			DSN: "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
		},
	})

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
