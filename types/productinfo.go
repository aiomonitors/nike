package types

type ProductInfo struct {
	Name           string   `json:"name"`
	SKU            string   `json:"sku"`
	Style          string   `json:"style"`
	Link           string   `json:"link"`
	Image          string   `json:"image"`
	AvailableSizes []string `json:"available"`
	Price          string   `json:"price"`
	Notification   string   `json:"notification"`
	DiscordSKUs    []string `json:"discord"`
	Exec           string   `json:"exec"`
}
