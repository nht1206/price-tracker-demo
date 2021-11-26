package main

import (
	"log"
	"os"

	"github.com/nht1206/pricetracker/config"
	app "github.com/nht1206/pricetracker/internal/app"
	"github.com/nht1206/pricetracker/internal/logger"
	"github.com/nht1206/pricetracker/static"
	"github.com/nht1206/pricetracker/system"
)

func main() {
	exitCode := static.APPLICATION_STATUS_SUCCESS
	defer func() {
		if exitCode != static.APPLICATION_STATUS_SUCCESS {
			os.Exit(exitCode)
		}
	}()
	cfg := &config.Config{
		AppName:                "Test",
		NumCrawlingGoroutines:  5,
		NumNotifyingGoroutines: 1,
		Log: &config.LogConfig{
			OutputPath: "./data/logs",
			FileName:   "test.log",
			Level:      "info",
		},
		DB: &config.DatabaseConfig{
			DSN: "root:toor@tcp(localhost:3306)/pricesubscriber",
		},
		Notifier: &config.NotifierConfig{
			Mail: &config.MailConfig{
				SMTPHost:       "smtp.gmail.com",
				SMTPPort:       "587",
				Sender:         "hp1t1nhy3u@gmail.com",
				SenderPassword: "01653374206",
			},
		},
	}

	err := logger.InitLogger(cfg.Log)
	if err != nil {
		exitCode = static.APPLICATION_STATUS_LOGGER_INIT_ERROR
		log.Printf("-----InitLogger-----\n err:%v", err)
		return
	}

	sysCtx, err := system.InitSystemContext(cfg)
	if err != nil {
		exitCode = static.APPLICATION_STATUS_CONTEXT_INIT_ERROR
		logger.Logger.Errorf("Failed to initialize app context. err: %v", err)
		return
	}

	app.StartApp(sysCtx)
}
