package get

import (
	"backlog/internal/api/response"
	"backlog/internal/api/task"
	"backlog/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Response struct {
	response.Response
	Tasks []task.Task `json:"tasks,omitempty"`
}

func OK(tasks []task.Task) Response {
	return Response{
		Tasks: tasks,
	}
}

type TaskGetter interface {
	GetTasks() ([]task.Task, error)
}

func New(log *slog.Logger, getter TaskGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(slog.String("request_id", middleware.GetReqID(r.Context())))

		tasks, err := getter.GetTasks()
		if err != nil {
			log.Error("couldn't get tasks", sl.Err(err))

			render.Status(r, 500)
			render.JSON(w, r, "couldn't get tasks")

			return
		}

		render.JSON(w, r, OK(tasks))
	}
}
