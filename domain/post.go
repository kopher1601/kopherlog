package domain

type PostCreate struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Post struct {
	Title   string
	Content string
}
