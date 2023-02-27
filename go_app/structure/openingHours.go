package structure

type OpeningHours struct {
	ID     int    `json:"id"`
	Day    int    `json:"day"`
	ShopID int    `json:"id_shop"`
	Open   string `json:"open"`
	Close  string `json:"close"`
}

type OpeningHoursStoreInterface interface {
	AddOpeningHours(opening_hours OpeningHours) (int, error)
}
