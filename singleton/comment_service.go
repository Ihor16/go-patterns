package singleton

import "patterns/singleton/database"

type CommentService struct {
}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (s CommentService) GetCommentById() {
	db := database.GetInstance()
	db.Query("comment")
}
