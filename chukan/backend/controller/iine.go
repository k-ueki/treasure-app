package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/service"
)

type Iine struct {
	db *sqlx.DB
}

func NewIine(db *sqlx.DB) *Iine {
	return &Iine{db: db}
}

func (i *Iine) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	iid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	req := &model.Iine{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	createIine := &model.Iine{
		IdeaID: iid,
		UserID: &user.ID,
	}
	fmt.Println(createIine)

	iineservice := service.NewIine(i.db)
	createdid, err := iineservice.Create(createIine)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	//
	// respParam := &model.ResponseCreateIdea{
	// 	ID:     id,
	// 	Title:  createIdea.Title,
	// 	Body:   createIdea.Body,
	// 	UserID: createIdea.UserID,
	// 	TagIDs: reqParam.TagIDs,
	// }

	return http.StatusCreated, createdid, nil
}
