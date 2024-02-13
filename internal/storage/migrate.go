package storage

//Based on https://github.com/evrone/go-clean-template

import (
	"log/slog"
	"time"

	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate/v4"
	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func Initialize(log *slog.Logger, storageURL string) error {
	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://storage/migrations", storageURL)
		if err == nil {
			break
		}

		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		return err
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "migrate: up error")
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Info("migrate: no change")
		return nil
	}

	log.Info("migrate: up success")
	return nil
}
