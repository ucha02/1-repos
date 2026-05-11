package myhttp

import (
	"TodoList/todo"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := NewErrDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := NewErrDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Text)
	if err := h.todoList.AddTask(todoTask); err != nil {
		writeErrorResponse(w, err)
		return
	}

	writeJSON(w, todoTask, http.StatusCreated)
}

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")

	task, err := h.todoList.GetTask(title)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}

	writeJSON(w, task, http.StatusOK)
}

func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	completedParam := query.Get("completed")

	tasks := h.todoList.GetAllTasks()

	if completedParam == "" {
		writeJSON(w, tasks, http.StatusOK)
		return
	}

	isDone, err := strconv.ParseBool(completedParam)
	if err != nil {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}

	completedFilter := make(map[string]*todo.Task)
	for k, v := range tasks {
		if v.IsDone == isDone {
			completedFilter[k] = v
		}
	}

	writeJSON(w, completedFilter, http.StatusOK)
}

func (h *HTTPHandlers) HandleStateChange(w http.ResponseWriter, r *http.Request) {
	var isDoneDTO CompleteTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&isDoneDTO); err != nil {
		errDTO := NewErrDTO(err)

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	title := chi.URLParam(r, "title")

	changedTask, err := h.todoList.TaskStateChange(isDoneDTO.IsDone, title)
	if err != nil {
		writeErrorResponse(w, err)
		return
	}

	writeJSON(w, changedTask, http.StatusOK)
}

func (h *HTTPHandlers) HandleDelTasks(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")

	if err := h.todoList.DelTask(title); err != nil {
		writeErrorResponse(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
