package entity

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	PassHash string `json:"pass_hash"`
}

func NewUser(username, passHash string) *User {
	return &User{
		Username: username,
		PassHash: passHash,
	}
}
