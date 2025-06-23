package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/paultustain/pomodoroTracker/m/v2/internal/database"
)

type Project struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	TimeSpent int32     `json:"time_spent"`
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

func (cfg *apiConfig) handlerGetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := cfg.db.GetProjects(r.Context())
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"failed to get projects: ",
			err,
		)
	}

	respondWithJSON(w, http.StatusOK, projects)
}

func (cfg *apiConfig) handlerUpdateTime(w http.ResponseWriter, r *http.Request) {
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

	project, err := cfg.db.GetProject(r.Context(), params.Name)
	if err != nil {
		respondWithError(w,
			http.StatusInternalServerError,
			"failed to find project: ",
			err,
		)
		return
	}

	updatedProject, err := cfg.db.UpdateTime(
		r.Context(),
		database.UpdateTimeParams{
			TimeSpent: project.TimeSpent + 1,
			Name:      project.Name,
		},
	)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to update time: ", err)
		return
	}

	respondWithJSON(
		w,
		http.StatusOK,
		Project{
			ID:        project.ID,
			UpdatedAt: project.UpdatedAt,
			CreatedAt: project.CreatedAt,
			Name:      project.Name,
			Completed: project.Completed,
			TimeSpent: updatedProject.TimeSpent,
		},
	)

}
