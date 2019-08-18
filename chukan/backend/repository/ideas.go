package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllIdeas(db *sqlx.DB) ([]model.Idea, error) {
	i := make([]model.Idea, 0)
	if err := db.Select(&i, `SELECT id, title, body, user_id FROM idea`); err != nil {
		return nil, err
	}
	return i, nil
}

func FindIdea(db *sqlx.DB, id int64) (*model.Idea, error) {
	i := model.Idea{}
	if err := db.Get(&i, `
SELECT id, title, body, user_id FROM idea WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &i, nil
}

func CreateIdea(db *sqlx.Tx, a *model.Idea) (sql.Result, error) {
	stmt, err := db.Prepare(`
 INSERT INTO idea (title, body, user_id) VALUES (?, ?, ?)
 `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body, a.UserID)
}

func UpdateIdea(db *sqlx.Tx, id int64, i *model.Idea) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE idea SET title = ?, body = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(i.Title, i.Body, id)
}

func DestroyIdea(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM idea WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}
