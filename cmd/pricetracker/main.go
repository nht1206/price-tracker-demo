package main

import (
	"context"
	"log"
	"os"

	"github.com/nht1206/pricetracker/config"
	app "github.com/nht1206/pricetracker/internal/app"
	pt_log "github.com/nht1206/pricetracker/internal/log"
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

	logger, err := pt_log.InitLogger(cfg.Log)
	if err != nil {
		exitCode = static.APPLICATION_STATUS_LOGGER_INIT_ERROR
		log.Printf("-----InitLogger-----\n err:%v", err)
		return
	}

	sysCtx, err := system.InitSystemContext(cfg)
	if err != nil {
		exitCode = static.APPLICATION_STATUS_CONTEXT_INIT_ERROR
		logger.Errorf("Failed to initialize app context, err: %v", err)
		return
	}

	ctx := context.Background()
	ctx = pt_log.WithLogger(ctx, logger)

	app.StartApp(ctx, sysCtx)
}
