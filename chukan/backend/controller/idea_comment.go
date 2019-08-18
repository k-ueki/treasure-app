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
	"github.com/voyagegroup/treasure-app/service"
)

type IdeaComment struct {
	dbx *sqlx.DB
}

func NewIdeaComment(dbx *sqlx.DB) *IdeaComment {
	return &IdeaComment{dbx: dbx}
}

func (ic *IdeaComment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["idea_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	reqParam := &model.CreateRequestIdeaComment{}
	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	createIdeaComment := &model.IdeaComment{
		UserID: user.ID,
		IdeaID: iid,
		Body:   reqParam.Body,
	}

	ideaCommentService := service.NewIdeaCommentService(ic.dbx)
	createdId, err := ideaCommentService.Create(createIdeaComment)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	createIdeaComment.ID = createdId

	return http.StatusCreated, createIdeaComment, nil
}
