package structure

type Reservation struct {
	ID        int    `json:"id_reservation"`
	ShopID    int    `json:"id_shop"`
	UserID    int    `json:"id_user"`
	BenefitID int    `json:"id_benefit"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Comment   string `json:"comment"`
}

type UpdateReservation struct {
	BenefitID int    `json:"id_benefit"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Comment   string `json:"comment"`
}

type ReservationRO struct {
	ShopID      int    `json:"id_shop"`
	BenefitID   int    `json:"id_benefit"`
	ShopName    string `json:"shop_name"`
	BenefitName string `json:"benefit_name"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Comment     string `json:"comment"`
}

type ExistingReservation struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

type ReservationStoreInterface interface {
	AddReservation(new_reservation Reservation, id_user int) (int, error)
	GetExistingReservationForPeriod(id_shop int) ([]ExistingReservation, error)
	GetReservationByUser(id_user int) ([]ReservationRO, error)
	UpdateReservation(id int, updated_reservation UpdateReservation) error
	DeleteReservation(id int) error
}
