package database

import (
	"back_project/structure"
	"database/sql"
)

func NewOpeningHoursStore(db *sql.DB) *OpeningHoursStore {
	return &OpeningHoursStore{
		db,
	}
}

type OpeningHoursStore struct {
	*sql.DB
}

func (openingHours_store *OpeningHoursStore) AddOpeningHours(new_openingHours structure.OpeningHours) (int, error) {
	res, err := openingHours_store.DB.Exec(
		"INSERT INTO openingHours (day, id_shop, open, close) VALUES (?, ?, ?, ?)",
		new_openingHours.Day,
		new_openingHours.ShopID,
		new_openingHours.Open,
		new_openingHours.Close)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
