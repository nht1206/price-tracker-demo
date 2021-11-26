package model

type Price struct {
	ID        uint64 `json:"id"`
	ProductID uint64 `json:"productId"`
	Price     string `json:"price"`
}

func (p Price) TableName() string {
	return "t_price"
}
