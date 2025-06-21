package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	err := cfg.db.DeleteProjects(r.Context())
	if err != nil {
		fmt.Printf("failed to delete projects: %w", err)
	}
	fmt.Println("Project database reset")
}
