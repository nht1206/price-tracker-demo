package service

import (
	"context"

	"github.com/nht1206/pricetracker/config"
	"github.com/nht1206/pricetracker/internal/logger"
	"github.com/nht1206/pricetracker/internal/model"
	"golang.org/x/sync/errgroup"
)

type notifyWorker struct {
	config *config.Config
}

func NewNotifyWorker(config *config.Config) *notifyWorker {
	return &notifyWorker{
		config: config,
	}
}

type notifyFuture struct {
	chTrackingResult chan<- model.TrackingResult
	wait             func() error
}

func (f *notifyFuture) Sink() chan<- model.TrackingResult {
	return f.chTrackingResult
}

func (f *notifyFuture) Wait() error {
	return f.wait()
}

func (w *notifyWorker) StartNotifying(ctx context.Context, cancel context.CancelFunc) *notifyFuture {
	fields := []interface{}{"Func", "PriceTracker.StartNotifying"}
	chTrackingResult := make(chan model.TrackingResult, w.config.NumNotifyingGoroutines)
	notifyingEg, notifyingCtx := errgroup.WithContext(ctx)
	for i := 0; i < w.config.NumNotifyingGoroutines; i++ {
		notifyingEg.Go(func() (notifyingErr error) {
			ctx := notifyingCtx
			defer func() {
				if notifyingErr != nil {
					cancel()
				}
			}()
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case v, ok := <-chTrackingResult:
					if !ok {
						return nil
					}
					logger.Logger.
						With(fields...).
						Infof("Success to inform the new price to user. productId: %v, oldPrice: %v, newPrice: %v", v.ProductId, v.OldPrice, v.NewPrice)
				}
			}
		})
	}

	future := &notifyFuture{
		chTrackingResult: chTrackingResult,
		wait: func() error {
			err := notifyingEg.Wait()
			if err != nil {
				return err
			}
			return nil
		},
	}

	return future
}
