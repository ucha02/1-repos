package myhttp

import (
	"TodoList/todo"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data any, HTTPstatus int) {
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(HTTPstatus)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response:", err)
		return
	}
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	errDTO := NewErrDTO(err)

	switch {
	case errors.Is(err, todo.ErrTaskNotFound):
		http.Error(w, errDTO.ToString(), http.StatusNotFound)

	case errors.Is(err, todo.ErrTaskAlreadyExists):
		http.Error(w, errDTO.ToString(), http.StatusConflict)

	default:
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
	}
}
