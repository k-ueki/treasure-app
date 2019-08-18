package model

type Iine struct {
	ID     int64  `db:"id" json:"id"`
	IdeaID int64  `db:"idea_id" json:"idea_id"`
	UserID *int64 `db:"user_id" json:"user_id"`
}
