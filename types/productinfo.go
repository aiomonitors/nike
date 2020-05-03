package types

type ProductInfo struct {
	Name           string   `json:"name"`
	SKU            string   `json:"sku"`
	Link           string   `json:"link"`
	AvailableSizes []string `json:"available"`
	Price          string   `json:"price"`
	Exec           string   `json:"exec"`
}
