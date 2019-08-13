package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
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
	fmt.Println("USER", user)

	if err := json.NewDecoder(r.Body).Decode(&newArticleComment); err != nil {
		fmt.Println("Decode")
		return http.StatusBadRequest, nil, err
	}

	fmt.Printf("%#v\n", newArticleComment)

	articleService := service.NewArticleCommentService(a.dbx)
	id, err := articleService.Create(newArticleComment)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newArticleComment.ID = id

	return http.StatusCreated, newArticleComment, nil
}

//func (a *Comment) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
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
//	reqComment := &model.Comment{}
//	if err := json.NewDecoder(r.Body).Decode(&reqComment); err != nil {
//		return http.StatusBadRequest, nil, err
//	}
//
//	articleService := service.NewComment(a.dbx)
//	err = articleService.Update(aid, reqComment)
//	if err != nil && errors.Cause(err) == sql.ErrNoRows {
//		return http.StatusNotFound, nil, err
//	} else if err != nil {
//		return http.StatusInternalServerError, nil, err
//	}
//
//	return http.StatusNoContent, nil, nil
//}

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
