package entity

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	PassHash string `json:"pass_hash"`
}
