package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/Faroukhamadi/northwind-api/ent"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// dependency injection
type server struct {
	router *gin.Engine
	db     *sql.DB
	orm    *ent.Client
	// l      *log.Logger
}

var s *server

func main() {
	ctx := context.Background()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer s.db.Close()
	defer s.orm.Close()
	countries, err := s.orm.Country.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first row: ", countries[0])

	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "cool",
		})
	})

	s.routes("hello world")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	go func() {
		log.Println("Starting server on port 8080")
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)

	sig := <-q
	log.Println("Got signal to shutdown server:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}

func run() error {
	db, err := sql.Open("pgx", "postgresql://postgres:faroukhamadi@127.0.0.1/world")
	if err != nil {
		return fmt.Errorf("open database %w", err)
	}
	// Create an ent.Driver from `db`
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))
	s = newServer(db, client, gin.Default())
	return nil
}

func newServer(db *sql.DB, orm *ent.Client, router *gin.Engine) *server {
	s := &server{
		router,
		db,
		orm,
	}
	return s
}

func (s *server) handleAPI() gin.HandlersChain {
	handler := s.router.Group("/api")

	handler.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"greeter": "hello",
		})
	})

	handler.GET("/bye", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"bye": "lataaa",
		})
	})
	return handler.Handlers
}

func (s *server) handleCountries() gin.HandlersChain {
	handler := s.router.Group("/countries")
	handler.GET("/tunisia", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":       "Tunisia",
			"population": "11.552.842",
		})
	})
	return handler.Handlers
}

func (s *server) routes(greeter string) {
	// implement this
	s.router.GET("/api", s.handleAPI()...)
	s.router.GET("/countries", s.handleCountries()...)
}
