package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

type ArticleComment struct {
	dbx *sqlx.DB
}

func NewArticleComment(dbx *sqlx.DB) *ArticleComment {
	return &ArticleComment{dbx: dbx}
}

func (a *ArticleComment) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	article_comments, err := repository.AllArticleComment(a.dbx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, article_comments, nil
}

//func (a *Comment) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
//	vars := mux.Vars(r)
//	id, ok := vars["id"]
//	if !ok {
//		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
//	}
//
//	aid, err := strconv.ParseInt(id, 10, 64)
//	if err != nil {
//		return http.StatusBadRequest, nil, err
//	}
//
//	article, err := repository.FindComment(a.dbx, aid)
//	if err != nil && err == sql.ErrNoRows {
//		return http.StatusNotFound, nil, err
//	} else if err != nil {
//		return http.StatusInternalServerError, nil, err
//	}
//
//	return http.StatusCreated, article, nil
//}

func (a *ArticleComment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	fmt.Println("article_id : ", id)

	newArticleComment := &model.ArticleComment{}

	contextUser, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		log.Print(err)
		return http.StatusBadRequest, nil, err
	}
	user, err := repository.GetUser(a.dbx, contextUser.FirebaseUID)

	if err != nil {
		fmt.Println("user")
		return http.StatusBadRequest, nil, err
	}
	fmt.Println(user.ID)
	newArticleComment.UserID = &user.ID
	fmt.Printf("%#v\n", user)

	if err := json.NewDecoder(r.Body).Decode(&newArticleComment); err != nil {
		fmt.Println("Decode")
		return http.StatusBadRequest, nil, err
	}

	intid, _ := strconv.ParseInt(id, 10, 64)
	newArticleComment.ArticleID = intid

	articlecommentService := service.NewArticleCommentService(a.dbx)
	ida, err := articlecommentService.Create(newArticleComment)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newArticleComment.ID = ida

	return http.StatusCreated, newArticleComment, nil
}

func (a *ArticleComment) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqArticleComment := &model.ArticleComment{}
	if err := json.NewDecoder(r.Body).Decode(&reqArticleComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	articlecommentService := service.NewArticleCommentService(a.dbx)
	err = articlecommentService.Update(aid, reqArticleComment)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

//func (a *Comment) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
//	vars := mux.Vars(r)
//	id, ok := vars["id"]
//	if !ok {
//		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
//	}
//
//	aid, err := strconv.ParseInt(id, 10, 64)
//	if err != nil {
//		return http.StatusBadRequest, nil, err
//	}
//
//	articleService := service.NewComment(a.dbx)
//	err = articleService.Destroy(aid)
//	if err != nil && errors.Cause(err) == sql.ErrNoRows {
//		return http.StatusNotFound, nil, err
//	} else if err != nil {
//		return http.StatusInternalServerError, nil, err
//	}
//
//	return http.StatusNoContent, nil, nil
//}
