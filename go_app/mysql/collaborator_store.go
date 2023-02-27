package database

import (
	"back_project/structure"
	"database/sql"
)

func NewCollaboratorStore(db *sql.DB) *CollaboratorStore {
	return &CollaboratorStore{
		db,
	}
}

type CollaboratorStore struct {
	*sql.DB
}

func (collaborator_store CollaboratorStore) GetCollaboratorByShop(id_shop int) ([]structure.Collaborator, error) {
	var collaborators []structure.Collaborator

	rows, err := collaborator_store.DB.Query("SELECT id_collaborator, name, phone, email FROM collaborators where id_shop = ? ", id_shop)
	if err != nil {
		return []structure.Collaborator{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collaborator structure.Collaborator
		if err = rows.Scan(
			&collaborator.ID,
			&collaborator.Name,
			&collaborator.Phone,
			&collaborator.Email); err != nil {
			return []structure.Collaborator{}, err
		}
		collaborators = append(collaborators, collaborator)
	}

	if err = rows.Err(); err != nil {
		return []structure.Collaborator{}, err
	}

	return collaborators, nil
}

func (collaborator_store CollaboratorStore) AddCollaborator(new_collaborator structure.Collaborator) (int, error) {
	res, err := collaborator_store.DB.Exec(
		"INSERT INTO collaborators (id_shop, name, phone, email) VALUES (?, ?, ?, ?)",
		new_collaborator.ShopID,
		new_collaborator.Name,
		new_collaborator.Phone,
		new_collaborator.Email)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (collaborator_store CollaboratorStore) DeleteCollaborator(id_collaborator int) error {
	_, err := collaborator_store.DB.Exec("DELETE FROM collaborators WHERE id_collaborator = ?", id_collaborator)
	if err != nil {
		return err
	}
	return nil
}

func (collaborator_store *CollaboratorStore) UpdateCollaborator(id_collaborator int, updated_collaborator structure.Collaborator) error {
	sqlStatement := `UPDATE collaborators SET 
		id_shop = ?,
		name = ?,
		phone = ?,
		email = ?
		WHERE id_collaborator = ?`

	_, err := collaborator_store.DB.Exec(sqlStatement,
		updated_collaborator.ShopID,
		updated_collaborator.Name,
		updated_collaborator.Phone,
		updated_collaborator.Email,
		id_collaborator,
	)
	if err != nil {
		return err
	}
	return nil
}
