package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllArticleComment(db *sqlx.DB) ([]model.ArticleComment, error) {
	a := make([]model.ArticleComment, 0)
	if err := db.Select(&a, `SELECT id, user_id, article_id, body FROM article`); err != nil {
		fmt.Println("RR")
		return nil, err
	}
	return a, nil
}

func FindArticleComment(db *sqlx.DB, id int64) (*model.ArticleComment, error) {
	a := model.ArticleComment{}
	if err := db.Get(&a, `SELECT body FROM article_comment WHERE id = ?`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateArticleComment(db *sqlx.Tx, a *model.ArticleComment) (sql.Result, error) {
	fmt.Println(a)
	stmt, err := db.Prepare(`INSERT INTO article_comment (user_id, article_id, body) VALUES (?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.UserID, a.ArticleID, a.Body)
}

func UpdateArticleComment(db *sqlx.Tx, id int64, a *model.ArticleComment) (sql.Result, error) {
	stmt, err := db.Prepare(`UPDATE article_comment SET body = ? WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Body, id)
}

//func DestroyArticle(db *sqlx.Tx, id int64) (sql.Result, error) {
//	stmt, err := db.Prepare(`
//DELETE FROM article WHERE id = ?
//`)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//	return stmt.Exec(id)
//}
