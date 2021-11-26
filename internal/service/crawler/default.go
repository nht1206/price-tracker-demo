package crawler

import "time"

type defaultCrawler struct {
}

func newDefaultCrawler() (Crawler, error) {
	return &defaultCrawler{}, nil
}

func (c *defaultCrawler) GetPrice(url string) (string, error) {
	time.Sleep(500 * time.Millisecond)
	return "30000", nil
}
