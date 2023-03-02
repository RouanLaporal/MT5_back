package structure

type Kind struct {
	ID   int    `json:"id_kind"`
	Name string `json:"name"`
}

type KindStoreInterface interface {
	GetAllKind() ([]Kind, error)
}

type NewShop struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ZipCode     string `json:"zip"`
	City        string `json:"city"`
	Lat         string `json:"lat"`
	Long        string `json:"long"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Description string `json:"description"`
	KindID      []int  `json:"id_kind"`
	UserID      int    `json:"id_user"`
}

type Shop struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ZipCode     string `json:"zip"`
	City        string `json:"city"`
	Lat         string `json:"lat"`
	Long        string `json:"long"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Description string `json:"description"`
	UserID      int    `json:"id_user"`
}

type ShopStoreInterface interface {
	// GetAllShopByKindAndPosition(id_type int) ([]Shop, error)
	GetAllShopByKindAndCity(id_kind int, city string) ([]Shop, error)
	GetAllShopByUser(id_user int) ([]Shop, error)
	AddShop(shop NewShop, id_user int) (int, error)
	DeleteShop(id int) error
	UpdateShop(id int, item Shop) error
}
