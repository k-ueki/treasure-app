package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/dbutil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Tag struct {
	db *sqlx.DB
}

func NewTag(db *sqlx.DB) *Tag {
	return &Tag{db}
}

func (t *Tag) Create(createTag *model.Tag) (int64, error) {

	var createdId int64
	if err := dbutil.TXHandler(t.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateTag(tx, createTag)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		createdId = id
		return err
	}); err != nil {
		return 0, err
	}
	return createdId, nil
}
