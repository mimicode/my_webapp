package tables

type User struct {
	UserId   int    `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func (User) TableName() string {
	return "user"
}
