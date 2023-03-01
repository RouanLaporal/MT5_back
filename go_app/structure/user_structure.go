package structure

type User struct {
	ID        int    `json:"id_user"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type UpdateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type AuthUser struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Role        string `json:"role"`
	TokenString string `json:"token"`
}
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Password struct {
	Password string `json:"password"`
}

type UserStoreInterface interface {
	GetUserByEmail(email string) (User, error)
	AddUser(item User) error
	DeleteUser(id int) error
	UpdateUser(id int, user UpdateUser) error
	UpdatePassword(email string, password string) error
}
