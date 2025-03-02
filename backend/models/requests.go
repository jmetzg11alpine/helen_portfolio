package models

type CreateBlogPostRequest struct {
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
	Content  string `json:"content"`
}
