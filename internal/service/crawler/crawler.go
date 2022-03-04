package crawler

import "errors"

var (
	ErrCrawlerNotFound = errors.New("crawler not found")
)

type CrawlerType string

const (
	Default CrawlerType = "default"
	Shopee  CrawlerType = "shopee"
)

var crawlers map[CrawlerType]Crawler

type Crawler interface {
	GetPrice(url string) (string, error)
}

func GetCrawler(crawlerType CrawlerType) (Crawler, error) {
	crawler, ok := crawlers[crawlerType]
	if ok {
		return crawler, nil
	}
	switch crawlerType {
	case Default:
		defaultCrawler, err := newDefaultCrawler()
		if err != nil {
			return nil, err
		}
		crawlers[Default] = defaultCrawler
		return defaultCrawler, nil
	case Shopee:
		shopeeCrawler := newShopeeCrawler()
		crawlers[Shopee] = shopeeCrawler
		return shopeeCrawler, nil
	default:
		return nil, ErrCrawlerNotFound
	}
}

func init() {
	crawlers = map[CrawlerType]Crawler{}
}
