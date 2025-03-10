package models

type CreateBlogPostRequest struct {
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Content  string `json:"content"`
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
