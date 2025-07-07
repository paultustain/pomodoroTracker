package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/paultustain/pomodoroTracker/m/v2/internal/database"
)

type Task struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	ProjectID uuid.UUID `json:"project_id"`
}

func (cfg *apiConfig) handlerCreateTask(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		ProjectID   uuid.UUID `json:"project_id"`
		Description string    `json:"task_description"`
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

	task, err := cfg.db.CreateTask(r.Context(), database.CreateTaskParams{
		Task: params.Description,
		ProjectID: uuid.NullUUID{
			UUID:  params.ProjectID,
			Valid: true,
		},
	})

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"failed to decode body: ",
			err,
		)
		return
	}
	respondWithJSON(
		w,
		http.StatusCreated,
		task,
	)
}

func (cfg *apiConfig) handlerGetProjectTasks(w http.ResponseWriter, r *http.Request) {

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

	tasks, err := cfg.db.GetProjectTasks(r.Context(), uuid.NullUUID{
		UUID:  projectID,
		Valid: true,
	})

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"failed to get tasks: ",
			err,
		)
		return
	}

	respondWithJSON(w, http.StatusOK, tasks)
}

func (cfg *apiConfig) handlerCompleteTask(w http.ResponseWriter, r *http.Request) {
	// Add task description to this bit
	// Add the description to the index/html and app.js files too

	taskIDString := r.PathValue("taskID")
	taskID, err := uuid.Parse(taskIDString)

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Invalid Task ID",
			err,
		)
		return
	}

	updatedTask, err := cfg.db.CompleteTask(r.Context(), taskID)
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"failed to update task to complete: ",
			err,
		)
	}

	respondWithJSON(w, http.StatusOK, updatedTask)

}

func (cfg *apiConfig) handlerDeleteTask(w http.ResponseWriter, r *http.Request) {

	taskIDString := r.PathValue("taskID")
	taskID, err := uuid.Parse(taskIDString)

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Invalid Task ID",
			err,
		)
		return
	}

	err = cfg.db.DeleteTask(r.Context(), taskID)

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Failed to delete task",
			err,
		)
		return
	}

}
