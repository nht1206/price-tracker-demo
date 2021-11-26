package logger

import (
	"os"
	"testing"

	"github.com/nht1206/pricetracker/config"
)

func TestInitLogger(t *testing.T) {
	err := InitLogger(&config.LogConfig{
		OutputPath: "",
		FileName:   "test.log",
		Level:      "info",
	})

	if err != nil {
		t.Errorf("failed to initialize logger. %v", err)
	}

	err = os.Remove("test.log")
	if err != nil {
		t.Errorf("failed to delete log file. %v", err)
	}
}
