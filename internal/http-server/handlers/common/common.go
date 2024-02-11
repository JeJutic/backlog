package common

import (
	"backlog/internal/api/response"
	"backlog/internal/api/task"
	"backlog/internal/lib/logger/sl"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

func DecodeTask(log *slog.Logger, w http.ResponseWriter, r *http.Request) (task.Task, bool) {
	var t task.Task

	err := render.DecodeJSON(r.Body, &t)
	if err != nil {
		log.Error("failed to decode request body", sl.Err(err))

		render.Status(r, 400)
		render.JSON(w, r, response.Error("failed to decode request body"))

		return t, false
	}

	log.Info("request body decoded", slog.Any("request", t))

	if err := validator.New().Struct(t); err != nil {
		log.Error("task validation failed", sl.Err(err))

		render.Status(r, 422)
		render.JSON(w, r, response.Error(err.Error()))

		return t, false
	}

	return t, true
}
