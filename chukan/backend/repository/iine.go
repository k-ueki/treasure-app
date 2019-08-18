package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

// func CreateIdeaTag(db *sqlx.Tx, ideaId int64, tagId int64) (sql.Result, error) {
// 	stmt, err := db.Prepare(`
// INSERT INTO idea_tag (idea_id, tag_id) VALUES (?, ?)
// `)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()
// 	return stmt.Exec(ideaId, tagId)
// }
//
func CreateIine(db *sqlx.Tx, createIine *model.Iine) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO iine (idea_id,user_id) VALUES (?,?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(createIine.IdeaID, createIine.UserID)
}

func FindIinesByIdeaId(db *sqlx.DB, ideaId int64) ([]model.Iine, error) {
	i := make([]model.Iine, 0)
	if err := db.Select(&i, `
SELECT id, idea_id, user_id FROM iine WHERE idea_id = ?
`, ideaId); err != nil {
		return nil, err
	}
	return i, nil
}

//
// func FindIdeasByTagID(db *sqlx.DB, tagId int64) ([]model.Idea, error) {
// 	i := make([]model.Idea, 0)
// 	if err := db.Select(&i, `
// SELECT idea.id as id,idea.title as title,idea.body as body,idea.user_id as user_id FROM idea_tag
// INNER JOIN idea ON idea.id = idea_tag.idea_id
// WHERE idea_tag.tag_id = ?
// `, tagId); err != nil {
// 		return nil, err
// 	}
// 	return i, nil
// }
