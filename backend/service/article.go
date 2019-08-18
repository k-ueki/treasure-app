package service

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Article struct {
	db *sqlx.DB
}

func NewArticle(db *sqlx.DB) *Article {
	return &Article{db}
}

func (a *Article) FindArticleDetail(id int64) (*model.ArticleDetail, error) {
	article, err := repository.FindArticle(a.db, id)
	if err != nil {
		return nil, err
	}
	tags, err := repository.FindArticleTagByArticleID(a.db, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	comments, err := repository.FindCommentsByArticleID(a.db, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	articleDetail := &model.ArticleDetail{
		Article:  *article,
		Tags:     tags,
		Comments: comments,
	}
	return articleDetail, nil
}

func (a *Article) Update(id int64, newArticle *model.Article) error {
	_, err := repository.FindArticle(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateArticle(tx, id, newArticle)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed article update transaction")
	}
	return nil
}

func (a *Article) Destroy(id int64) error {
	_, err := repository.FindArticle(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.DestroyArticle(tx, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed article delete transaction")
	}
	return nil
}

func (a *Article) Create(createArticle *model.Article, tagIds []int64) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateArticle(tx, createArticle)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		for _, tagId := range tagIds {
			_, err = repository.CreateArticleTag(tx, id, tagId)
			if err != nil {
				return err
			}
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
