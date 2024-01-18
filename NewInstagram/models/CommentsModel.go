package models

type CommentModel struct {
	Id int
	UserID int
	PostID int
	Content string
}