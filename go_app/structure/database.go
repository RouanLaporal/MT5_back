package structure

type Token struct {
	IDUser      int    `json:"id_user"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}
