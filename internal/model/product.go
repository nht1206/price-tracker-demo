package model

type Product struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (p Product) TableName() string {
	return "t_product"
}
