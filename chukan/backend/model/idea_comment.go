package model

type CreateRequestIdeaComment struct {
	Body string `json:"body"`
}

type IdeaComment struct {
	ID     int64  `db:"id" json:"id"`
	Body   string `db:"body" json:"body"`
	IdeaID int64  `db:"idea_id" json:"idea_id"`
	UserID int64  `db:"user_id" json:"user_id"`
}
