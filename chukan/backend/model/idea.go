package model

type RequestCreateIdea struct {
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	TagIDs []int64 `json:"tag_ids"`
}

type ResponseCreateIdea struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	UserID *int64  `json:"user_id"`
	TagIDs []int64 `json:"tag_ids"`
}

type IdeaDetail struct {
	Idea
	Comments []IdeaComment `json:"comments"`
	Tags     []Tag         `json:"tags"`
	Iine     []Iine        `json:"iine"`
}
type Idea struct {
	ID     int64  `db:"id" json:"id"`
	Title  string `db:"title" json:"title"`
	Body   string `db:"body" json:"body"`
	UserID *int64 `db:"user_id" json:"user_id"`
}
