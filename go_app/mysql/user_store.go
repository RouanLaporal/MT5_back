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
	err := s.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Phone, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserStore) AddUser(item structure.User) (int, error) {
	var id int
	hashPassword, _ := helper.HashPassword(item.Password)

	item.Password = hashPassword
	err := s.QueryRow("INSERT INTO users (name, phone, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id", item.Name, item.Phone, item.Email, item.Password, item.Role).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *UserStore) DeleteUser(id int) error {
	_, err := s.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// func (s *UserStore) UpdateUser(id int) error {
// 	_, err := s.Exec("UPDATE users SET name = $1, phone = $2, email = $3, password = $4 WHERE id = $5", id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
