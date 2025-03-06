package domain

type PostCreate struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Post struct {
	ID      int
	Title   string
	Content string
}
