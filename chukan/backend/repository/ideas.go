package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllIdeas(db *sqlx.DB) ([]model.Idea, error) {
	i := make([]model.Idea, 0)
	if err := db.Select(&i, `SELECT id, title, body, user_id FROM ideas`); err != nil {
		return nil, err
	}
	return i, nil
}

// func FindIdeas(db *sqlx.DB, id int64) (*model.Idea, error) {
// 	a := model.Ideas{}
// 	if err := db.Get(&a, `
// SELECT id, title, body, user_id FROM article WHERE id = ?
// `, id); err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }
//
// func CreateIdeas(db *sqlx.Tx, a *model.Idea) (sql.Result, error) {
// 	stmt, err := db.Prepare(`
// INSERT INTO article (title, body, user_id) VALUES (?, ?, ?)
// `)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()
// 	return stmt.Exec(a.Title, a.Body, a.UserID)
// }
//
// func UpdateIdeas(db *sqlx.Tx, id int64, a *model.Idea) (sql.Result, error) {
// 	stmt, err := db.Prepare(`
// UPDATE article SET title = ?, body = ? WHERE id = ?
// `)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()
// 	return stmt.Exec(a.Title, a.Body, id)
// }
//
// func DestroyIdeas(db *sqlx.Tx, id int64) (sql.Result, error) {
// 	stmt, err := db.Prepare(`
// DELETE FROM article WHERE id = ?
// `)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()
// 	return stmt.Exec(id)
// }
