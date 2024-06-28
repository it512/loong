package loong

import (
	"context"
	"log/slog"
)

type Config struct {
	Store        Store
	EventHandler EventHandler
	templates    TemplateGetter
	IoConnector  IoConnector

	Txer Txer

	ctx context.Context

	Logger *slog.Logger
}

type Option func(*Config)

func SetTxer(tx Txer) Option {
	return func(e *Config) {
		e.Txer = tx
	}
}

func SetIoConnector(sc IoConnector) Option {
	return func(e *Config) {
		e.IoConnector = sc
	}
}

func SetContext(ctx context.Context) Option {
	return func(e *Config) {
		e.ctx = ctx
	}
}

func SetStore(s Store) Option {
	return func(e *Config) {
		e.Store = s
	}
}

func SetEventHandler(eh EventHandler) Option {
	return func(e *Config) {
		e.EventHandler = eh
	}
}

func SetLogger(logger *slog.Logger) Option {
	return func(e *Config) {
		e.Logger = logger
	}
}
