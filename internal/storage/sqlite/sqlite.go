package sqlite

import (
	"backlog/internal/api/task"
	"backlog/internal/storage"
	"database/sql"
	"log/slog"

	"github.com/pkg/errors"
	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func New(log *slog.Logger, storagePath string) (*Storage, error) {
	storageURL := "sqlite://" + storagePath
	err := storage.Initialize(log, storageURL)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't initialize storage: %s", storagePath)
	}

	db, err := sql.Open("sqlite", storagePath)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't connect to db: %s", storagePath)
	}

	return &Storage{db}, nil
}

func (s *Storage) CreateTask(text string) error {
	stmt, err := s.db.Prepare(
		"INSERT INTO tasks(text, status_id) VALUES " +
			"(?, (SELECT id FROM task_status WHERE name = 'To do'))",
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't prepare insertion of task with text: %q", text)
	}

	_, err = stmt.Exec(text)
	return errors.Wrapf(err, "task with text: %q wasn't inserted", text)
}

func (s *Storage) GetTasks() ([]task.Task, error) {
	stmt, err := s.db.Prepare(
		"SELECT tasks.id, text, task_status.name FROM tasks JOIN task_status " +
			"ON status_id = task_status.id",
	)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't prepare statement to get tasks")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get tasks")
	}

	var tasks []task.Task

	for rows.Next() {
		var t task.Task
		err := rows.Scan(&t.Id, &t.Text, &t.Status)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't scan rows")
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (s *Storage) MoveTask(id int64, status string) error {
	stmt, err := s.db.Prepare(
		"UPDATE tasks SET status_id = (SELECT id FROM task_status WHERE name = ?) " +
			"WHERE id = ?",
	)
	if err != nil {
		return errors.Wrapf(err, "couldn't prepare update of task with id: %v", id)
	}

	_, err = stmt.Exec(status, id)
	return errors.Wrapf(err, "couldn't update task with id: %v", id)
}
