package user

type User struct {
	UserID       int    `json:"userId"`
	Fname        string `json:"fname"`
	Lname        string `json:"lname"`
	Age          int    `json:"age"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"passwordHash"`
}
