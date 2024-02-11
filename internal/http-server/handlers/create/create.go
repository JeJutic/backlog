package create

import (
	"backlog/internal/api/response"
	"backlog/internal/http-server/handlers/common"
	"backlog/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type TaskCreator interface {
	CreateTask(text string) error
}

func New(log *slog.Logger, creator TaskCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(slog.String("request_id", middleware.GetReqID(r.Context())))

		t, ok := common.DecodeTask(log, w, r)
		if !ok {
			return
		}

		if t.Text == "" {
			log.Error("task text was empty")

			render.Status(r, 422)
			render.JSON(w, r, response.Error("task text can't be empty"))

			return
		}

		err := creator.CreateTask(t.Text)
		if err != nil {
			log.Error("task creation failed", sl.Err(err))

			render.Status(r, 500)
			render.JSON(w, r, response.Error("task creation failed"))

			return
		}
	}
}
