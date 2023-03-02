package structure

type Benefit struct {
	IDBenefit   int    `json:"id_benefit"`
	IDShop      int    `json:"id_shop"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Price       string `json:"price"`
}

type BenefitRO struct {
	IDBenefit   int    `json:"id_benefit"`
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
