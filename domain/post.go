package domain

import "math"

type PostCreate struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostSearchParams struct {
	Page int `form:"page" validate:"required,numeric,gte=0"`
	Size int `form:"size" validate:"required,numeric,gte=0,lte=100"`
}

type PostSearch struct {
	Page int
	Size int
}

func (p *PostSearch) Offset() int {
	return int(math.Max(1, float64(p.Page))-1) * p.Size
}

func (p *PostSearch) Limit() int {
	const maxSize = 100
	return int(math.Min(maxSize, float64(p.Size)))
}

type Post struct {
	ID      int
	Title   string
	Content string
}
