package models
type UserModel struct {
	ID int
	Firsname string
	Lastname string
	Post []PostModel
}