package main

import (
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/Faroukhamadi/northwind-api/api"
)

// // dependency injection
// type server struct {
// 	router *gin.Engine
// 	db     *sql.DB
// 	orm    *ent.Client
// 	// l      *log.Logger
// }

func main() {
	// s := api.Init()
	server := api.Init()
	log.Println("Listening on port", server.Addr[1:])
	server.ListenAndServe()

	// srv := &http.Server{
	// 	Addr:    ":9090",
	// 	Handler: s,
	// }

	// ctx := context.Background()
	// if err := run(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err)
	// 	os.Exit(1)
	// }
	// defer s.db.Close()
	// defer s.orm.Close()
	// countries, err := s.orm.Country.Query().All(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("first row: ", countries[0])

	// s.router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"data": "cool",
	// 	})
	// })

	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: s.router,
	// }

	// go func() {
	// 	log.Println("Starting server on port 8080")
	// 	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalf("Error starting server: %s\n", err)
	// 	}
	// }()

	// q := make(chan os.Signal, 1)
	// signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)

	// sig := <-q
	// log.Println("Got signal to shutdown server:", sig)

	// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer cancel()

	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server forced to shutdown:", err)
	// }

	// log.Println("Server exiting")

}

// func run() error {
// 	db, err := sql.Open("pgx", "postgresql://postgres:faroukhamadi@127.0.0.1/world")
// 	if err != nil {
// 		return fmt.Errorf("open database %w", err)
// 	}
// 	// Create an ent.Driver from `db`
// 	drv := entsql.OpenDB(dialect.Postgres, db)
// 	client := ent.NewClient(ent.Driver(drv))
// 	s = newServer(db, client, gin.Default())
// 	return nil
// }

// func newServer(db *sql.DB, orm *ent.Client, router *gin.Engine) *server {
// 	s := &server{
// 		router,
// 		db,
// 		orm,
// 	}
// 	return s
// }
