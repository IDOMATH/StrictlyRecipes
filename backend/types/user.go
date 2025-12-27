package types

type User struct {
	Id           int
	Email        string
	PasswordHash string
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
