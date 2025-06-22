package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
}

func (cfg *apiConfig) handlerProjectCreate(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Params{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"failed to decode body: ",
			err,
		)
		return
	}

	project, err := cfg.db.CreateProject(r.Context(), params.Name)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"failed to create project: ",
			err,
		)
		return
	}
	respondWithJSON(
		w, http.StatusCreated, Project{
			ID:        project.ID,
			CreatedAt: project.CreatedAt,
			UpdatedAt: project.UpdatedAt,
			Name:      project.Name,
			Completed: project.Completed,
		},
	)
}

func (cfg *apiConfig) HandlerGetProjects(w http.ResponseWriter, r *http.Request) {
	videos, err := cfg.db.GetProjects(r.Context())
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"failed to get videos: ",
			err,
		)
	}

	respondWithJSON(w, http.StatusOK, videos)
}
