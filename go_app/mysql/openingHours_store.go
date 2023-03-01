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
		"INSERT INTO openings (id_day, id_shop, open, close) VALUES (?, ?, ?, ?)",
		new_openingHours.DayID,
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

func (openingHours_store *OpeningHoursStore) GetOpeningHoursByShop(id_shop int) (structure.ShowOpening, error) {
	var show_opening structure.ShowOpening
	rows := openingHours_store.DB.QueryRow("SELECT open, close, id_day FROM openings WHERE id_shop = ?", id_shop)
	switch err := rows.Scan(&show_opening.Open, &show_opening.Close, &show_opening.DayID); err {
	case sql.ErrNoRows:
		return show_opening, err
	case nil:
		return show_opening, nil
	default:
		return show_opening, err
	}
}

func (openingHours_store *OpeningHoursStore) UpdateOpeningHours(id int, updated_openingHours structure.OpeningHours) error {
	sqlStatement := `UPDATE openings SET 
	open = ?,
	close = ?
	WHERE id = ?`

	_, err := openingHours_store.DB.Exec(sqlStatement,
		updated_openingHours.Open,
		updated_openingHours.Close,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (openingHours_store *OpeningHoursStore) DeleteOpeningHours(id int) error {
	sqlStatement := `DELETE FROM openings WHERE id = ?`
	_, err := openingHours_store.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
