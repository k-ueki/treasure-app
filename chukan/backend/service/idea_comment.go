package service

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/dbutil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type IdeaComment struct {
	db *sqlx.DB
}

func NewIdeaCommentService(db *sqlx.DB) *IdeaComment {
	return &IdeaComment{db}
}

func (ic *IdeaComment) Create(createIdeaComment *model.IdeaComment) (int64, error) {
	fmt.Println(createIdeaComment.IdeaID)
	_, err := repository.FindIdea(ic.db, createIdeaComment.IdeaID)
	if err != nil {
		fmt.Println("OK")
		return 0, err
	}
	var createdId int64
	if err := dbutil.TXHandler(ic.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateIdeaComment(tx, createIdeaComment)
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
