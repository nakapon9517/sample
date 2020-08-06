package handler

// User struct
type User struct {
	Id       string `gorm:primary_key`
	Name     string
	Email    string
	Password string
}
