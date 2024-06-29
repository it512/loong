package loong

import (
	"context"
	"log/slog"
)

type Config struct {
	Store        Storer
	EventHandler EventHandler
	templates    TemplateGetter
	IoConnector  IoConnector

	Txer Txer

	Context context.Context

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
		e.Context = ctx
	}
}

func SetStorer(s Storer) Option {
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
