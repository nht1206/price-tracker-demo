package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/nht1206/pricetracker/config"
	"github.com/nht1206/pricetracker/internal/logger"
	"github.com/nht1206/pricetracker/internal/model"
	"github.com/nht1206/pricetracker/internal/repository"
	"github.com/nht1206/pricetracker/internal/service/crawler"
	"github.com/nht1206/pricetracker/static"
	"golang.org/x/sync/errgroup"
)

type priceTracker struct {
	config      *config.Config
	productRepo repository.ProductRepository
}

func NewPriceTracker(config *config.Config,
	productRepo repository.ProductRepository) *priceTracker {
	return &priceTracker{
		config:      config,
		productRepo: productRepo,
	}
}

type priceTrackerFuture struct {
	wait func() error
}

func (f *priceTrackerFuture) Wait() error {
	return f.wait()
}

func (w *priceTracker) StartTracking(ctx context.Context, cancel context.CancelFunc,
	products []model.Product, chNotify chan<- model.TrackingResult) *priceTrackerFuture {

	fields := []interface{}{"Func", "PriceTracker.StartTracking"}
	chInputProduct := make(chan model.Product, w.config.NumCrawlingGoroutines)
	crawlingEg, crawlingCtx := errgroup.WithContext(ctx)
	for i := 0; i < w.config.NumCrawlingGoroutines; i++ {
		crawlingEg.Go(func() (crawlingErr error) {
			ctx := crawlingCtx
			defer func() {
				if crawlingErr != nil {
					cancel()
				}
			}()
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case v, ok := <-chInputProduct:
					if !ok {
						return nil
					}
					result, err := w.trackPrice(v)
					if err != nil {
						return err
					}
					logger.Logger.
						With(fields...).
						Infof("Success to track product price. productId: %v, oldPrice: %v, newPrice: %v", result.ProductId, result.OldPrice, result.NewPrice)
					oldPrice, err := strconv.ParseInt(result.OldPrice, 10, 64)
					if err != nil {
						return err
					}
					newPrice, err := strconv.ParseInt(result.OldPrice, 10, 64)
					if err != nil {
						return err
					}
					if newPrice < oldPrice {
						select {
						case <-ctx.Done():
							return ctx.Err()
						case chNotify <- *result:
						}
					}
				}
			}
		})
	}

	pushDataEg, pushDataCtx := errgroup.WithContext(ctx)
	pushDataEg.Go(func() (pushDataErr error) {
		ctx := pushDataCtx
		defer func() {
			if pushDataErr != nil {
				cancel()
			}
		}()
		for _, p := range products {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case chInputProduct <- p:
			}
		}

		return nil
	})

	future := &priceTrackerFuture{
		wait: func() error {
			err := pushDataEg.Wait()
			if err != nil {
				return err
			}
			close(chInputProduct)

			err = crawlingEg.Wait()
			if err != nil {
				return err
			}
			close(chNotify)

			return nil
		},
	}

	return future
}

func (w *priceTracker) trackPrice(product model.Product) (*model.TrackingResult, error) {
	successFlg := false
	rowAffected, err := w.productRepo.LockProductToTrackPrice(product.ID)
	if err != nil || rowAffected != static.MINIMUM_ROW_AFFECTED {
		return nil, fmt.Errorf("can not lock the product. productId: %v, err: %v, rowAffected: %v", product.ID, err, rowAffected)
	}
	defer func() {
		if !successFlg {
			w.productRepo.UpdateProductStatusToFailed(product.ID)
		}
	}()
	oldPrice, err := w.productRepo.GetProductPrice(product.ID)
	if err != nil {
		return nil, err
	}

	crawlerType := getCrawlerType(product.URL)

	defaultCrawler, err := crawler.GetCrawler(crawlerType)
	if err != nil {
		return nil, err
	}
	newPrice, err := defaultCrawler.GetPrice(product.URL)
	if err != nil {
		return nil, err
	}

	if newPrice != oldPrice.Price {
		rowAffected, err = w.productRepo.UpdateProductPrice(product.ID, newPrice)
		if err != nil || rowAffected != static.MINIMUM_ROW_AFFECTED {
			return nil, fmt.Errorf("can not update price for the product. productId: %v, err: %v, rowAffected: %v", product.ID, err, rowAffected)
		}
	}

	rowAffected, err = w.productRepo.UnlockProduct(product.ID)
	if err != nil || rowAffected != static.MINIMUM_ROW_AFFECTED {
		return nil, fmt.Errorf("can not unlock the product. productId: %v, err: %v, rowAffected: %v", product.ID, err, rowAffected)
	}

	successFlg = true

	return &model.TrackingResult{
		ProductId: product.ID,
		OldPrice:  oldPrice.Price,
		NewPrice:  newPrice,
	}, nil
}

func getCrawlerType(url string) crawler.CrawlerType {
	if strings.HasPrefix(url, "https://shopee.vn/") {
		return crawler.Shopee
	}
	return crawler.Default
}
