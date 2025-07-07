package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/paultustain/pomodoroTracker/m/v2/internal/database"
)

type Project struct {
	ID            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	TimeSpent     int32     `json:"time_spent"`
	TimeLimitType string    `json:"time_limit_type"`
	TimeLimit     int32     `json:"time_limit"`
	Completed     bool      `json:"completed"`
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

	project, err := cfg.db.CreateProject(r.Context(), database.CreateProjectParams{
		Name:          params.Name,
		TimeLimitType: "Tracking Only",
	})

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
			TimeSpent: project.TimeSpent,
			Completed: project.Completed,
		},
	)
}

func (cfg *apiConfig) handlerGetProjects(w http.ResponseWriter, r *http.Request) {
	// Change this pathValue for the ID?
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
		ID        uuid.UUID `json:"id"`
		TimeAdded int32     `json:"time"`
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
	project, err := cfg.db.GetProject(r.Context(), params.ID)
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
			TimeSpent: project.TimeSpent + params.TimeAdded,
			ID:        project.ID,
		},
	)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to update time: ", err)
		return
	}

	respondWithJSON(
		w,
		http.StatusOK,
		updatedProject,
	)
}

func (cfg *apiConfig) handlerDeleteProject(w http.ResponseWriter, r *http.Request) {
	projectIDString := r.PathValue("projectID")
	projectID, err := uuid.Parse(projectIDString)

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Invalid Project ID",
			err,
		)
		return
	}

	err = cfg.db.DeleteProject(r.Context(), projectID)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Failed to delete project",
			err,
		)
	}

}
