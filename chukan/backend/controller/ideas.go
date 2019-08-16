package controller

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/repository"
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

// func (a *Ideas) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
// 	vars := mux.Vars(r)
// 	id, ok := vars["id"]
// 	if !ok {
// 		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
// 	}
//
// 	aid, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		return http.StatusBadRequest, nil, err
// 	}
//
// 	ideaservice := service.NewIdeas(a.db)
// 	articleDetail, err := ideaservice.FindIdeas(aid)
// 	if err != nil && err == sql.ErrNoRows {
// 		return http.StatusNotFound, nil, err
// 	} else if err != nil {
// 		return http.StatusInternalServerError, nil, err
// 	}
// 	return http.StatusCreated, articleDetail, nil
// }
//
// func (a *Ideas) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
// 	reqParam := &model.RequestCreateIdeas{}
// 	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
// 		return http.StatusBadRequest, nil, err
// 	}
//
// 	user, err := httputil.GetUserFromContext(r.Context())
// 	if err != nil {
// 		return http.StatusInternalServerError, nil, err
// 	}
//
// 	createIdeas := &model.Ideas{
// 		Title:  reqParam.Title,
// 		Body:   reqParam.Body,
// 		UserID: &user.ID,
// 	}
//
// 	ideaservice := service.NewIdeas(a.db)
// 	id, err := ideaservice.Create(createIdeas, reqParam.TagIDs)
// 	if err != nil {
// 		return http.StatusInternalServerError, nil, err
// 	}
//
// 	respParam := &model.ResponseCreateIdeas{
// 		ID:     id,
// 		Title:  createIdeas.Title,
// 		Body:   createIdeas.Body,
// 		UserID: createIdeas.UserID,
// 		TagIDs: reqParam.TagIDs,
// 	}
//
// 	return http.StatusCreated, respParam, nil
// }
//
// // func (a *Ideas) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
// // 	vars := mux.Vars(r)
// // 	id, ok := vars["id"]
// // 	if !ok {
// // 		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
// // 	}
// //
// // 	aid, err := strconv.ParseInt(id, 10, 64)
// // 	if err != nil {
// // 		return http.StatusBadRequest, nil, err
// // 	}
// //
// // 	reqParam := &model.Ideas{}
// // 	if err := json.NewDecoder(r.Body).Decode(&reqParam); err != nil {
// // 		return http.StatusBadRequest, nil, err
// // 	}
// //
// // 	ideaservice := service.NewIdeas(a.db)
// // 	err = ideaservice.Update(aid, reqParam)
// // 	if err != nil && errors.Cause(err) == sql.ErrNoRows {
// // 		return http.StatusNotFound, nil, err
// // 	} else if err != nil {
// // 		return http.StatusInternalServerError, nil, err
// // 	}
// //
// // 	return http.StatusNoContent, nil, nil
// // }
// //
// // func (a *Ideas) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
// // 	vars := mux.Vars(r)
// // 	id, ok := vars["id"]
// // 	if !ok {
// // 		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
// // 	}
// //
// // 	aid, err := strconv.ParseInt(id, 10, 64)
// // 	if err != nil {
// // 		return http.StatusBadRequest, nil, err
// // 	}
// //
// // 	ideaservice := service.NewIdeas(a.db)
// // 	err = ideaservice.Destroy(aid)
// // 	if err != nil && errors.Cause(err) == sql.ErrNoRows {
// // 		return http.StatusNotFound, nil, err
// // 	} else if err != nil {
// // 		return http.StatusInternalServerError, nil, err
// // 	}
// //
// // 	return http.StatusNoContent, nil, nil
// // }
