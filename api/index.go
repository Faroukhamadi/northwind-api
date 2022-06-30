package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Faroukhamadi/northwind-api/ent"
	"github.com/gorilla/mux"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"

	_ "github.com/Faroukhamadi/northwind-api/cmd/server/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// dependecy injection
type Server struct {
	l      *log.Logger
	router *mux.Router
	db     *sql.DB
	orm    *ent.Client
}

func newServer(l *log.Logger, r *mux.Router, db *sql.DB, orm *ent.Client) *Server {
	return &Server{l, r, db, orm}
}

func Init() (srv *http.Server) {
	s, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	getR := s.router.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/countries", s.GetOne).
		Queries("id", "{id:[A-Z]+}")
	getR.HandleFunc("/countries", s.GetAll)

	postR := s.router.Methods(http.MethodPut).Subrouter()
	postR.HandleFunc("/countries", s.UpdateOne).
		Queries("id", "{id:[A-Z]+}")

	s.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9090/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	srv = &http.Server{
		Addr:    ":9090",
		Handler: s,
	}

	return srv

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func run() (s *Server, err error) {
	db, err := sql.Open("pgx", "postgresql://postgres:faroukhamadi@127.0.0.1/world")
	if err != nil {
		return
	}
	// Create an ent.Driver from `db`
	drv := entsql.OpenDB(dialect.Postgres, db)
	// Initialize ent client
	client := ent.NewClient(ent.Driver(drv))
	// Create custom logger
	l := log.New(os.Stdout, "country-api", log.LstdFlags)
	// Create gorilla router
	sm := mux.NewRouter()

	s = newServer(l, sm, db, client)

	return
}
