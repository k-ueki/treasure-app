package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func CreateIdeaTag(db *sqlx.Tx, ideaId int64, tagId int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO idea_tag (idea_id, tag_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(ideaId, tagId)
}

func CreateTag(db *sqlx.Tx, createTag *model.Tag) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO tag (name) VALUES (?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(createTag.Name)
}

func FindIdeaTagByIdeaID(db *sqlx.DB, ideaId int64) ([]model.Tag, error) {
	t := make([]model.Tag, 0)
	if err := db.Select(&t, `
SELECT tag.id as id, tag.name as name FROM idea_tag 
INNER JOIN tag ON tag.id = idea_tag.tag_id
WHERE idea_tag.idea_id = ?
`, ideaId); err != nil {
		return nil, err
	}
	return t, nil
}

func FindIdeasByTagID(db *sqlx.DB, tagId int64) ([]model.Idea, error) {
	i := make([]model.Idea, 0)
	if err := db.Select(&i, `
SELECT idea.id as id,idea.title as title,idea.body as body,idea.user_id as user_id FROM idea_tag
INNER JOIN idea ON idea.id = idea_tag.idea_id
WHERE idea_tag.tag_id = ?
`, tagId); err != nil {
		return nil, err
	}
	return i, nil
}
