package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/paultustain/pomodoroTracker/m/v2/internal/config"
)

const ROOTDIR = "./app"
const PORT = "8080"

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Println(cfg.DBURL)
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(ROOTDIR))

	mux.Handle("/", fs)
	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + PORT,
	}
	fmt.Printf("File serving on port: localhost:8080")

	server.ListenAndServe()
}
