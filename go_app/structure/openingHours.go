package structure

type OpeningHours struct {
	ID     int    `json:"id"`
	DayID  int    `json:"id_day"`
	ShopID int    `json:"id_shop"`
	Open   string `json:"open"`
	Close  string `json:"close"`
}

type ShowOpening struct {
	Open  string `json:"open"`
	Close string `json:"close"`
	DayID string `json:"day"`
}

type OpeningHoursStoreInterface interface {
	AddOpeningHours(opening_hours OpeningHours) (int, error)
	GetOpeningHoursByShop(id_shop int) ([]ShowOpening, error)
	UpdateOpeningHours(id int, opening_hours OpeningHours) error
	DeleteOpeningHours(id int) error
}
