package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func CreateIdeaComment(db *sqlx.Tx, ic *model.IdeaComment) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO idea_comment (body, idea_id, user_id) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(ic.Body, ic.IdeaID, ic.UserID)
}

func FindCommentsByIdeaID(db *sqlx.DB, ideaId int64) ([]model.IdeaComment, error) {
	ic := make([]model.IdeaComment, 0)
	if err := db.Select(&ic, `
SELECT id, idea_id, user_id, body FROM idea_comment WHERE idea_id = ?
`, ideaId); err != nil {
		return nil, err
	}
	return ic, nil
}
