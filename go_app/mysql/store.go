package database

import (
	"back_project/structure"
	"database/sql"
)

func CreateStore(db *sql.DB) *Store {
	return &Store{
		NewUserStore(db),
		NewKindStore(db),
		NewShopStore(db),
		NewBenefitStore(db),
		NewReviewStore(db),
		NewCollaboratorStore(db),
		NewOpeningHoursStore(db),
		NewReservationStore(db),
	}
}

type Store struct {
	structure.UserStoreInterface
	structure.KindStoreInterface
	structure.ShopStoreInterface
	structure.BenefitStoreInterface
	structure.ReviewStoreInterface
	structure.CollaboratorStoreInterface
	structure.OpeningHoursStoreInterface
	structure.ReservationStoreInterface
}
