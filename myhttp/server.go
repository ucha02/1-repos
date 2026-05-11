package myhttp

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := chi.NewRouter()
	router.Post("/tasks", s.httpHandlers.HandleCreateTask)
	router.Get("/tasks/{title}", s.httpHandlers.HandleGetTask)
	router.Get("/tasks", s.httpHandlers.HandleGetAllTasks)
	router.Patch("/tasks/{title}", s.httpHandlers.HandleStateChange)
	router.Delete("/tasks/{title}", s.httpHandlers.HandleDelTasks)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
