package structure

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserStoreInterface interface {
	GetUserByEmail(email string) (User, error)
	AddUser(item User) (int, error)
	DeleteUser(id int) error
	// UpdateUser(id int) error // TODO : update user
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticationStoreInterface interface {
	SignIn(email string) (Authentication, error)
	// SignUp(item User) (int, error)
}
