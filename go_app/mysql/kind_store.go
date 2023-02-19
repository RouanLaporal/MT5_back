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

func (s *KindStore) GetAllKind() (structure.Kind, error) {
	var kind structure.Kind
	rows := s.DB.QueryRow("SELECT id_kind, name FROM kind")
	switch err := rows.Scan(&kind.ID, &kind.Name); err {
	case sql.ErrNoRows:
		return kind, err
	case nil:
		return kind, nil
	default:
		return kind, err
	}
}
