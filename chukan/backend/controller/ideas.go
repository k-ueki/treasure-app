package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Ideas struct {
	db *sqlx.DB
}

func NewIdeas(db *sqlx.DB) *Ideas {
	return &Ideas{db: db}
}

func (a *Ideas) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	ideas, err := repository.AllIdeas(a.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, ideas, nil
}

func (i *Ideas) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ideaservice := service.NewIdea(i.db)
	ideaDetail, err := ideaservice.FindIdeaDetail(iid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusCreated, ideaDetail, nil
}

func (i *Ideas) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	reqParam := &model.RequestCreateIdea{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	createIdea := &model.Idea{
		Title:  reqParam.Title,
		Body:   reqParam.Body,
		UserID: &user.ID,
	}

	ideaservice := service.NewIdea(i.db)
	id, err := ideaservice.Create(createIdea, reqParam.TagIDs)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	respParam := &model.ResponseCreateIdea{
		ID:     id,
		Title:  createIdea.Title,
		Body:   createIdea.Body,
		UserID: createIdea.UserID,
		TagIDs: reqParam.TagIDs,
	}

	return http.StatusCreated, respParam, nil
}

func (i *Ideas) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqParam := &model.Idea{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	ideaservice := service.NewIdea(i.db)
	err = ideaservice.Update(iid, reqParam)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (i *Ideas) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ideaservice := service.NewIdea(i.db)
	err = ideaservice.Destroy(iid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}
