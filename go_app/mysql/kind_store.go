package database

import (
	"back_project/structure"
	"database/sql"
)

func NewKindStore(db *sql.DB) *KindStore {
	return &KindStore{
		db,
	}
}

type KindStore struct {
	*sql.DB
}

func (kind_store *KindStore) GetAllKind() ([]structure.Kind, error) {
	var kinds []structure.Kind
	rows, err := kind_store.DB.Query("SELECT id_kind, name FROM kinds")
	if err != nil {
		return []structure.Kind{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var kind structure.Kind
		if err = rows.Scan(
			&kind.ID,
			&kind.Name); err != nil {
			return []structure.Kind{}, err
		}
		kinds = append(kinds, kind)
	}
	if err = rows.Err(); err != nil {
		return []structure.Kind{}, err
	}

	return kinds, nil
}
