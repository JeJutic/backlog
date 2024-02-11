package update

import (
	"backlog/internal/api/response"
	"backlog/internal/http-server/handlers/common"
	"backlog/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type TaskMover interface {
	MoveTask(id int64, status string) error
}

func New(log *slog.Logger, mover TaskMover) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(slog.String("request_id", middleware.GetReqID(r.Context())))

		t, ok := common.DecodeTask(log, w, r)
		if !ok {
			return
		}

		err := mover.MoveTask(t.Id, t.Status)
		if err != nil {
			log.Error("task or status not found", sl.Err(err))

			render.Status(r, 422)
			render.JSON(w, r, response.Error("task or status not found"))

			return
		}
	}
}
