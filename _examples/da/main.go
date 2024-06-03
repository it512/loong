package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/it512/loong"
	"github.com/it512/loong/mongox"
	"github.com/it512/loong/web/gql"

	"github.com/it512/da/internal/io"
)

func main() {
	loong.Logo()

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}

	// DATABASE_URL := os.Getenv("DATABASE_URL")
	MONGODB_URI := os.Getenv("MONGODB_URI")
	SERVER_ADDR := os.Getenv("SERVER_ADDR")

	// db := loong.Must(pgx.OpenDB(DATABASE_URL))

	eng := loong.NewEngine(
		"loong-da",
		mongox.MongoStore(MONGODB_URI),
		loong.FileTemplates("./bpmn/", "*.bpmn"),
		loong.SetIoConnector(new(io.Io)),
	)

	if err := eng.Run(); err != nil {
		log.Fatal(err)
	}

	mux := chi.NewMux()
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Mount("/gql", gql.Playground("GraphQL playground", "/gql/query"))
	mux.Mount("/gql/query", gql.New(eng))

	log.Printf("SERVER_ADDR is -> %s", SERVER_ADDR)

	if err := http.ListenAndServe(SERVER_ADDR, mux); err != nil {
		log.Fatal(err)
	}
}
