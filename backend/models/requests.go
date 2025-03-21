package models

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateBlogPostRequest struct {
	Title    string `json:"title" binding:"required"`
	SubTitle string `json:"sub_title" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type CreateBlogCommentRequest struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	BlogID  uint   `json:"blogID"`
}

type CommentIDRequest struct {
	ID uint `json:"id" binding:"required"`
}
