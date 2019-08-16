package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Idea struct {
	db *sqlx.DB
}

func NewIdea(db *sqlx.DB) *Idea {
	return &Idea{db}
}

func (i *Idea) FindIdeaDetail(id int64) (*model.IdeaDetail, error) {
	idea, err := repository.FindIdea(i.db, id)
	if err != nil {
		return nil, err
	}
	// tags, err := repository.FindIdeaTagByIdeaID(i.db, id)
	// if err != nil && err != sql.ErrNoRows {
	// 	return nil, err
	// }
	// comments, err := repository.FindCommentsByIdeaID(i.db, id)
	// if err != nil && err != sql.ErrNoRows {
	// 	return nil, err
	// }
	ideaDetail := &model.IdeaDetail{
		Idea: *idea,
		// Tags: tags,
		// Comments: comments,
	}
	return ideaDetail, nil
}

func (i *Idea) Update(id int64, newIdea *model.Idea) error {
	_, err := repository.FindIdea(i.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := dbutil.TXHandler(i.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateIdea(tx, id, newIdea)
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

func (i *Idea) Destroy(id int64) error {
	_, err := repository.FindIdea(i.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := dbutil.TXHandler(i.db, func(tx *sqlx.Tx) error {
		_, err := repository.DestroyIdea(tx, id)
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

func (i *Idea) Create(createIdea *model.Idea, tagIds []int64) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(i.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateIdea(tx, createIdea)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		// for _, tagId := range tagIds {
		// 	_, err = repository.CreateArticleTag(tx, id, tagId)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
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
