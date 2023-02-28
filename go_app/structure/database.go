package structure

type User struct {
	ID        int    `json:"id_user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type UserStoreInterface interface {
	GetUserByEmail(email string) (User, error)
	AddUser(item User) (int, error)
	DeleteUser(id int) error
	UpdateUser(id int, user User) error
}

type AuthUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Role        string `json:"role"`
	TokenString string `json:"token"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
