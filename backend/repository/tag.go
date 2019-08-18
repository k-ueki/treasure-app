package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func CreateArticleTag(db *sqlx.Tx, articleId int64, tagId int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article_tag (article_id, tag_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(articleId, tagId)
}

func FindArticleTagByArticleID(db *sqlx.DB, articleId int64) ([]model.Tag, error) {
	t := make([]model.Tag, 0)
	if err := db.Select(&t, `
SELECT tag.id as id, tag.name as name FROM article_tag 
INNER JOIN tag ON tag.id = article_tag.tag_id
WHERE article_tag.article_id = ?
`, articleId); err != nil {
		return nil, err
	}
	return t, nil
}
