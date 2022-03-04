package pricetracker

import (
	"context"

	"github.com/nht1206/pricetracker/internal/log"
	"github.com/nht1206/pricetracker/internal/service"
	"github.com/nht1206/pricetracker/static"
	"github.com/nht1206/pricetracker/system"
)

func StartApp(ctx context.Context, sysCtx *system.Context) {
	logger := log.FromContext(ctx)
	logger = logger.With("app", sysCtx.Config.AppName)
	ctx = log.WithLogger(ctx, logger)

	logger.Info("App starts")

	targetTrackingProducts, err := sysCtx.Dao.FindTargetTrackingProduct()
	if err != nil {
		logger.Errorf("Failed at FindTargetTrackingProduct. err: %v", err)
		return
	}

	if len(targetTrackingProducts) == static.NO_TARGET {
		logger.Info("No target tracking products.")
	} else {
		ctx, cancel := context.WithCancel(ctx)

		notifyingFuture := service.NewNotifyWorker(sysCtx.Config, sysCtx.Dao, sysCtx.NotifierFactory).
			StartNotifying(ctx, cancel)

		priceCrawlingFuture := service.NewPriceTracker(sysCtx.Config, sysCtx.Dao).
			StartTracking(ctx, cancel, targetTrackingProducts, notifyingFuture.Sink())

		err = priceCrawlingFuture.Wait()
		if err != nil {
			logger.Errorf("Failed to crawl prices for the products. err: %v", err)
			return
		}

		err = notifyingFuture.Wait()
		if err != nil {
			logger.Errorf("Failed to notify prices to users. err: %v", err)
			return
		}
	}

	logger.Info("App ends.")
}
