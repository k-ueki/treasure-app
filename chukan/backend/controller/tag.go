package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/service"
)

type Tag struct {
	db *sqlx.DB
}

func NewTag(db *sqlx.DB) *Tag {
	return &Tag{db: db}
}

func (t *Tag) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	req := &model.Tag{}
	///ここまでしなくて良い??
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagservice := service.NewTag(t.db)
	id, err := tagservice.Create(req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	req.ID = id

	return http.StatusCreated, req, nil

}
