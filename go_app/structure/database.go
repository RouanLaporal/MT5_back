package structure

type Token struct {
	IDUser      int    `json:"id_user"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
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
	KindID      string `json:"id_kind"`
	UserID      int    `json:"id_user"`
}

type Kind struct {
	ID   int    `json:"id_kind"`
	Name string `json:"name"`
}

type Benefit struct {
	IDBenefit   int    `json:"id_benefit"`
	IDShop      int    `json:"id_shop"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Price       string `json:"price"`
}

type Review struct {
	IDReview int    `json:"id_review"`
	IDShop   int    `json:"id_shop"`
	IDUser   int    `json:"id_user"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
}

type KindStoreInterface interface {
	GetAllKind() (Kind, error)
}

type ShopStoreInterface interface {
	// GetAllShopByKindAndPosition(id_type int) ([]Shop, error)
	GetAllShopByKindAndCity(id_kind int, city string) ([]Shop, error)
	GetAllShopByUser(id_user int) ([]Shop, error)
	AddShop(shop Shop) (int, error)
	DeleteShop(id int) error
	UpdateShop(id int, item Shop) error
}

type BenefitStoreInterface interface {
	GetBenefitByShop(id_shop int) ([]Benefit, error)
	AddBenefit(benefit Benefit) (int, error)
	UpdateBenefit(id_benefit int, item Benefit) error
	DeleteBenefit(id_benefit int) error
}

type ReviewStoreInterface interface {
	GetReviewByShop(id_shop int) ([]Review, error)
	AddReview(review Review) (int, error)
	UpdateReview(id_review int, item Review) error
	DeleteReview(id_review int) error
}
