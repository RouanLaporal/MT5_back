package structure

type Review struct {
	IDReview int    `json:"id_review"`
	IDShop   int    `json:"id_shop"`
	IDUser   int    `json:"id_user"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
}

type ReviewRO struct {
	IDReview int    `json:"id_review"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
}

type ReviewStoreInterface interface {
	GetReviewByShop(id_shop int) ([]Review, error)
	AddReview(review Review, id_user int) (int, error)
	UpdateReview(id_review int, item Review) error
	DeleteReview(id_review int) error
}
