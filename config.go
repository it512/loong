package loong

import (
	"context"
	"log/slog"
)

type Config struct {
	store     Store
	eh        EventHandler
	templates TemplateGetter
	connector IoConnector

	ctx context.Context

	logger *slog.Logger

	queueSize uint
}

type Option func(*Config)

func SetIoConnector(sc IoConnector) Option {
	return func(e *Config) {
		e.connector = sc
	}
}

func SetContext(ctx context.Context) Option {
	return func(e *Config) {
		e.ctx = ctx
	}
}

func SetStore(s Store) Option {
	return func(e *Config) {
		e.store = s
	}
}

func SetEventHandler(eh EventHandler) Option {
	return func(e *Config) {
		e.eh = eh
	}
}

func SetLogger(logger *slog.Logger) Option {
	return func(e *Config) {
		e.logger = logger
	}
}
