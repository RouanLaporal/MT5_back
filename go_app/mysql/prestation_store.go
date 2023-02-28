package database

import (
	"back_project/structure"
	"database/sql"
)

func NewBenefitStore(db *sql.DB) *BenefitStore {
	return &BenefitStore{
		db,
	}
}

type BenefitStore struct {
	*sql.DB
}

func (s *BenefitStore) AddBenefit(item structure.Benefit) (int, error) {

	res, err := s.DB.Exec("INSERT INTO benefits (id_shop, name, description, duration, price) VALUES (?, ?, ?, ?, ?)", item.IDShop, item.Name, item.Description, item.Duration, item.Price)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *BenefitStore) GetBenefitByShop(id_shop int) ([]structure.Benefit, error) {
	var benefits []structure.Benefit

	rows, err := s.DB.Query("SELECT id_benefit, id_shop, name, description, duration, price from benefits WHERE id_shop = ?", id_shop)
	if err != nil {
		return []structure.Benefit{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var benefit structure.Benefit
		if err = rows.Scan(
			&benefit.IDBenefit,
			&benefit.IDShop,
			&benefit.Name,
			&benefit.Description,
			&benefit.Duration,
			&benefit.Price,
		); err != nil {
			return []structure.Benefit{}, err
		}
		benefits = append(benefits, benefit)
	}

	if err = rows.Err(); err != nil {
		return []structure.Benefit{}, err
	}

	return benefits, nil
}

func (s *BenefitStore) UpdateBenefit(id_benefit int, updated_benefit structure.Benefit) error {
	sqlStatement := `UPDATE benefits SET
		name = ?,
		description = ?,
		duration = ?,
		price = ?
	WHERE id_benefit = ?
	`
	_, err := s.DB.Exec(sqlStatement, updated_benefit.Name, updated_benefit.Description, updated_benefit.Duration, updated_benefit.Price, id_benefit)

	if err != nil {
		return err
	}
	return nil
}

func (s *BenefitStore) DeleteBenefit(id_benefit int) error {
	_, err := s.DB.Exec("DELETE from benefits WHERE id_benefit = ?", id_benefit)

	if err != nil {
		return err
	}
	return nil
}
