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
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = logger.InitLogger(cfg.Log)
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
