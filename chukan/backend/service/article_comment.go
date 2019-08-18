package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type ArticleComment struct {
	db *sqlx.DB
}

func NewArticleCommentService(db *sqlx.DB) *ArticleComment {
	return &ArticleComment{db}
}

func (ac *ArticleComment) Create(createArticleComment *model.ArticleComment) (int64, error) {
	_, err := repository.FindArticle(ac.db, createArticleComment.ArticleID)
	if err != nil {
		return 0, err
	}
	var createdId int64
	if err := dbutil.TXHandler(ac.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateArticleComment(tx, createArticleComment)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed article insert transaction")
	}
	return createdId, nil
}
