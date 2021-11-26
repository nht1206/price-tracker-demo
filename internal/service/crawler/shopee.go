package crawler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	ShoppeProductInfoAPI = `https://shopee.vn/api/v4/item/get?shopid=%v&itemid=%v`
)

type shopeeCrawler struct {
}

func newShopeeCrawler() Crawler {
	return &shopeeCrawler{}
}

func (c *shopeeCrawler) GetPrice(url string) (string, error) {
	shopId, productId := getIdFromURL(url)

	return getPrice(shopId, productId)
}

func getPrice(shopId, productId string) (string, error) {
	resp, err := http.Get(fmt.Sprintf(ShoppeProductInfoAPI, shopId, productId))
	if err != nil {
		return "", err
	}

	item := make(map[string]json.RawMessage)
	content, _ := io.ReadAll(resp.Body)

	json.Unmarshal(content, &item)
	data := make(map[string]json.RawMessage)
	json.Unmarshal(item["data"], &data)

	tempPrice := string(data["price"])

	price := ""

	if tempPrice != "" && len(tempPrice) > 5 {
		price = tempPrice[:len(tempPrice)-5]
	} else {
		return "", fmt.Errorf("can not get price of the shopee product")
	}

	return price, nil
}

func getIdFromURL(url string) (string, string) {
	splittedUrl := strings.Split(url, "/")

	productId := splittedUrl[len(splittedUrl)-1]
	shopId := splittedUrl[len(splittedUrl)-2]

	return shopId, productId
}
