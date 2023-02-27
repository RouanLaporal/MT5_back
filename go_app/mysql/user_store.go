package database

import (
	"back_project/helper"
	"back_project/structure"
	"database/sql"
)

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

type UserStore struct {
	*sql.DB
}

func (s *UserStore) GetUserByEmail(email string) (structure.User, error) {
	var user structure.User
	rows := s.DB.QueryRow("SELECT id_user, first_name, last_name, phone, email, password, role FROM users WHERE email = ?", email)
	switch err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Phone, &user.Email, &user.Password, &user.Role); err {
	case sql.ErrNoRows:
		return user, err
	case nil:
		return user, nil
	default:
		return user, err
	}
}

func (s *UserStore) AddUser(item structure.User) (int, error) {
	hashPassword, _ := helper.HashPassword(item.Password)

	item.Password = hashPassword
	res, err := s.DB.Exec("INSERT INTO users (first_name, last_name, phone, email, password, role) VALUES (?, ?, ?, ?, ?, ?)", item.FirstName, item.LastName, item.Phone, item.Email, item.Password, item.Role)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *UserStore) DeleteUser(id int) error {
	_, err := s.DB.Exec("DELETE FROM users WHERE id_user = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) UpdateUser(id_user int, updateUser structure.User) error {
	hashPassword, _ := helper.HashPassword(updateUser.Password)
	updateUser.Password = hashPassword

	sqlStatement := ` UPDATE users SET 
		first_name = ?,
		last_name = ?, 
		phone = ?, 
		email = ?, 
		password = ?
	WHERE id_user = ?`

	_, err := s.Exec(sqlStatement, updateUser.FirstName, updateUser.LastName, updateUser.Phone, updateUser.Email, updateUser.Password, id_user)
	if err != nil {
		return err
	}
	return nil
}
