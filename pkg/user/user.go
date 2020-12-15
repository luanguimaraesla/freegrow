package user

type User struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
