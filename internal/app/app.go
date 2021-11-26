package pricetracker

import (
	"context"

	"github.com/nht1206/pricetracker/internal/logger"
	"github.com/nht1206/pricetracker/internal/service"
	"github.com/nht1206/pricetracker/static"
	"github.com/nht1206/pricetracker/system"
)

func StartApp(sysCtx *system.Context) {
	fields := []interface{}{
		"App",
		sysCtx.Config.AppName,
	}
	logger.Logger.
		With(fields...).Info("App starts")

	targetTrackingProducts, err := sysCtx.Dao.FindTargetTrackingProduct()
	if err != nil {
		logger.Logger.
			With(fields...).Errorf("Failed at FindTargetTrackingProduct. err: %v", err)
		return
	}

	if len(targetTrackingProducts) == static.NO_TARGET {
		logger.Logger.
			With(fields...).Info("No target tracking products.")
	} else {
		ctx, cancel := context.WithCancel(context.Background())

		notifyingFuture := service.NewNotifyWorker(sysCtx.Config, sysCtx.Dao).
			StartNotifying(ctx, cancel)

		priceCrawlingFuture := service.NewPriceTracker(sysCtx.Config, sysCtx.Dao).
			StartTracking(ctx, cancel, targetTrackingProducts, notifyingFuture.Sink())

		err = priceCrawlingFuture.Wait()
		if err != nil {
			logger.Logger.
				With(fields...).Errorf("Failed to crawl prices for the products. err: %v", err)
			return
		}

		err = notifyingFuture.Wait()
		if err != nil {
			logger.Logger.
				With(fields...).Errorf("Failed to notify prices to users. err: %v", err)
			return
		}
	}

	logger.Logger.
		With(fields...).Info("App ends.")
}
