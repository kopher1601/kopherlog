package domain

type PostCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
