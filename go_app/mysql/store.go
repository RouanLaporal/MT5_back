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
		NewCollaboratorStore(db),
		NewOpeningHoursStore(db),
		// NewBenefitStore(db),
		// NewReviewStore(db),
		// NewReservationStore(db),
	}
}

type Store struct {
	structure.UserStoreInterface
	structure.KindStoreInterface
	structure.ShopStoreInterface
	structure.CollaboratorStoreInterface
	structure.OpeningHoursStoreInterface
	// structure.BenefitStoreInterface
	// structure.ReviewStoreInterface
	// structure.ReservationStoreInterface
}
