package database

import (
	"back_project/structure"
	"database/sql"
)

func NewReservationStore(db *sql.DB) *ReservationStore {
	return &ReservationStore{
		db,
	}
}

type ReservationStore struct {
	*sql.DB
}

func (reservation_store *ReservationStore) AddReservation(new_reservation structure.Reservation, id_user int) (int, error) {
	res, err := reservation_store.DB.Exec("INSERT INTO reservations (id_shop, id_user, id_benefit, date, time, comment) VALUES (?, ?, ?, ?, ?, ?)",
		new_reservation.ShopID,
		id_user,
		new_reservation.BenefitID,
		new_reservation.Date,
		new_reservation.Time,
		new_reservation.Comment)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (reservation_store *ReservationStore) GetExistingReservationForPeriod(id_shop int) ([]structure.ExistingReservation, error) {
	var reservations []structure.ExistingReservation
	rows, err := reservation_store.DB.Query("SELECT date, time FROM reservations WHERE id_shop = ? AND  date BETWEEN DATE(NOW())  AND DATE_ADD(DATE(NOW()), INTERVAL 90 day)", id_shop)
	if err != nil {
		return []structure.ExistingReservation{}, err
	}
	for rows.Next() {
		var reservation structure.ExistingReservation
		if err = rows.Scan(
			&reservation.Date,
			&reservation.Time); err != nil {
			return []structure.ExistingReservation{}, err
		}
		reservations = append(reservations, reservation)
	}
	if err = rows.Err(); err != nil {
		return []structure.ExistingReservation{}, err

	}
	defer rows.Close()

	return reservations, nil
}

func (reservation_store *ReservationStore) GetReservationByUser(id_user int) ([]structure.ReservationRO, error) {
	var reservations []structure.ReservationRO

	rows, err := reservation_store.DB.Query("SELECT reservations.id_shop, shops.name, reservations.id_benefit, benefits.name, date, time, comment FROM reservations RIGHT JOIN shops ON reservations.id_shop = shops.id_shop LEFT JOIN benefits ON reservations.id_benefit = benefits.id_benefit  where reservations.id_user = ?", id_user)
	if err != nil {
		return []structure.ReservationRO{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var reservation structure.ReservationRO
		if err = rows.Scan(
			&reservation.ShopID,
			&reservation.ShopName,
			&reservation.BenefitID,
			&reservation.BenefitName,
			&reservation.Date,
			&reservation.Time,
			&reservation.Comment,
		); err != nil {
			return []structure.ReservationRO{}, err
		}
		reservations = append(reservations, reservation)
	}

	if err = rows.Err(); err != nil {
		return []structure.ReservationRO{}, err
	}

	return reservations, nil
}

func (reservation_store *ReservationStore) UpdateReservation(id int, updated_reservation structure.UpdateReservation) error {
	sqlStatement := `UPDATE reservations SET 
	id_benefit = ?,
	date = ?,
	time = ?,
	comment = ? 
	WHERE id_reservation = ?`

	_, err := reservation_store.DB.Exec(sqlStatement,
		updated_reservation.BenefitID,
		updated_reservation.Date,
		updated_reservation.Time,
		updated_reservation.Comment,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (reservation_store *ReservationStore) DeleteReservation(id int) error {
	_, err := reservation_store.DB.Exec("DELETE FROM reservations WHERE id_reservation = ?", id)
	if err != nil {
		return err
	}
	return nil
}
