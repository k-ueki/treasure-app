package service

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/dbutil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
)

type Iine struct {
	db *sqlx.DB
}

func NewIine(db *sqlx.DB) *Iine {
	return &Iine{db}
}

func (i *Iine) Create(createIine *model.Iine) (int64, error) {

	var createdId int64
	if err := dbutil.TXHandler(i.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateIine(tx, createIine)
		if err != nil {
			fmt.Println("JJ")
			fmt.Println(err)
			return err
		}
		fmt.Println("OK")
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
