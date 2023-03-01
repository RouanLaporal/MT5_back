package structure

type Kind struct {
	ID   int    `json:"id_kind"`
	Name string `json:"name"`
}

type KindStoreInterface interface {
	GetAllKind() ([]Kind, error)
}

type Benefit struct {
	IDBenefit   int    `json:"id_benefit"`
	IDShop      int    `json:"id_shop"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Price       string `json:"price"`
}

type BenefitStoreInterface interface {
	GetBenefitByShop(id_shop int) ([]Benefit, error)
	AddBenefit(benefit Benefit) (int, error)
	UpdateBenefit(id_benefit int, item Benefit) error
	DeleteBenefit(id_benefit int) error
}

type Review struct {
	IDReview int    `json:"id_review"`
	IDShop   int    `json:"id_shop"`
	IDUser   int    `json:"id_user"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
}

type ReviewStoreInterface interface {
	GetReviewByShop(id_shop int) ([]Review, error)
	AddReview(review Review) (int, error)
	UpdateReview(id_review int, item Review) error
	DeleteReview(id_review int) error
}
