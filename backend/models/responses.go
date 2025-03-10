package models

type BlogPostPreview struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	SubTitle     string `json:"sub_title"`
	CreatedAt    int64  `json:"created_at"`
	CommentCount int64  `json:"comment_count"`
}

type UnapprovedComments struct {
	BlogComment
	Title    string `json:"title"`
	SubTitle string `json:"sub_title"`
}
