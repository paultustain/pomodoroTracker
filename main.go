package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/paultustain/pomodoroTracker/m/v2/internal/config"
	"github.com/paultustain/pomodoroTracker/m/v2/internal/database"
)

const ROOTDIR = "./app"
const PORT = "8080"

type apiConfig struct {
	db *database.Queries
}

func main() {

	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("failed to get filepath: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	apiCfg := apiConfig{
		db: dbQueries,
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(ROOTDIR))

	// base api functions
	mux.Handle("/", fs)
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	// Project functions
	mux.HandleFunc("POST /api/createProject", apiCfg.handlerProjectCreate)
	mux.HandleFunc("POST /api/createTask", apiCfg.handlerCreateTask)
	mux.HandleFunc("GET /api/getProjects", apiCfg.handlerGetProjects)
	mux.HandleFunc("POST /api/updateTime", apiCfg.handlerUpdateTime)
	mux.HandleFunc("GET /api/getTasks/{projectID}", apiCfg.handlerGetProjectTasks)
	mux.HandleFunc("POST /api/completeTask/{taskID}", apiCfg.handlerCompleteTask)
	mux.HandleFunc("DELETE /api/deleteProject/{projectID}", apiCfg.handlerDeleteProject)
	mux.HandleFunc("DELETE /api/deleteTask/{taskID}", apiCfg.handlerDeleteTask)

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + PORT,
	}
	fmt.Println("File serving on port: localhost:8080")

	server.ListenAndServe()
}
