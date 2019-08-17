package server

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"firebase.google.com/go/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"github.com/voyagegroup/treasure-app/controller"
	db2 "github.com/voyagegroup/treasure-app/db"
	"github.com/voyagegroup/treasure-app/firebase"
	"github.com/voyagegroup/treasure-app/middleware"
	"github.com/voyagegroup/treasure-app/sample"
)

type Server struct {
	db         *sqlx.DB
	router     *mux.Router
	authClient *auth.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {
	authClient, err := firebase.InitAuthClient()
	if err != nil {
		log.Fatalf("failed init auth client. %s", err)
	}
	s.authClient = authClient

	db := db2.NewDB(datasource)
	dbcon, err := db.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuth(s.authClient, s.db)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	authChain := commonChain.Append(
		authMiddleware.Handler,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/public").Handler(commonChain.Then(sample.NewPublicHandler()))
	r.Methods(http.MethodGet).Path("/private").Handler(authChain.Then(sample.NewPrivateHandler(s.db)))

	ideasController := controller.NewIdeas(s.db)
	r.Methods(http.MethodPost).Path("/ideas").Handler(authChain.Then(AppHandler{ideasController.Create}))
	r.Methods(http.MethodPut).Path("/ideas/{id}").Handler(authChain.Then(AppHandler{ideasController.Update}))
	r.Methods(http.MethodDelete).Path("/ideas/{id}").Handler(authChain.Then(AppHandler{ideasController.Destroy}))
	r.Methods(http.MethodGet).Path("/ideas").Handler(commonChain.Then(AppHandler{ideasController.Index}))
	r.Methods(http.MethodGet).Path("/ideas/{id}").Handler(commonChain.Then(AppHandler{ideasController.Show}))
	r.Methods(http.MethodGet).Path("/ideas/tag/{tag_id}").Handler(commonChain.Then(AppHandler{ideasController.TagSearch}))

	iineController := controller.NewIine(s.db)
	r.Methods(http.MethodPost).Path("/ideas/{id}/iine").Handler(authChain.Then(AppHandler{iineController.Create}))

	ideaCommentController := controller.NewIdeaComment(s.db)
	r.Methods(http.MethodPost).Path("/ideas/{idea_id}/comments").Handler(authChain.Then(AppHandler{ideaCommentController.Create}))

	tagController := controller.NewTag(s.db)
	r.Methods(http.MethodPost).Path("/tag").Handler(authChain.Then(AppHandler{tagController.Create}))

	r.PathPrefix("").Handler(commonChain.Then(http.StripPrefix("/img", http.FileServer(http.Dir("./img")))))
	return r
}
