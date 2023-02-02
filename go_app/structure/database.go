package structure

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStoreInterface interface {
	GetUserByEmail(email string) (User, error)
	AddUser(item User) (int, error)
	DeleteUser(id int) error
	// UpdateUser(id int) error // TODO : update user
}
