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

type ShopRO struct {
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
}

type ShopsNear struct {
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
	Kind string  `json:"kind"`
}

type ShopsNearReturn struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	ZipCode         string  `json:"zip"`
	City            string  `json:"city"`
	Lat             float64 `json:"lat"`
	Long            float64 `json:"long"`
	Country         string  `json:"country"`
	Phone           string  `json:"phone"`
	Email           string  `json:"email"`
	Description     string  `json:"description"`
	DistanceInMiles float64 `json:"distanceInMiles"`
}

type ShopStoreInterface interface {
	GetAllShopByKindAndCity(id_kind int, city string) ([]Shop, error)
	GetAllShopByUser(id_user int) ([]Shop, error)
	AddShop(shop NewShop, id_user int) (int, error)
	DeleteShop(id int) error
	UpdateShop(id int, item Shop) error
	GetAllShopNear(lat float64, long float64, kind string) ([]ShopsNearReturn, error)
	GetShopById(id int) (Shop, error)
}
