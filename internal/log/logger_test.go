package log

import (
	"testing"

	"github.com/nht1206/pricetracker/config"
)

func TestInitLogger(t *testing.T) {
	_, err := InitLogger(&config.LogConfig{
		Level: "info",
	})

	if err != nil {
		t.Errorf("failed to initialize logger. %v", err)
	}
}
