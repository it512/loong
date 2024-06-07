package loong

import (
	"context"
	"log"
	"log/slog"
	"runtime/debug"
)

type liquid struct {
	loop chan Activity
	c    context.Context

	ech chan Activity
	eh  EventHandler

	// cmdCh chan Cmd

	size uint

	logger *slog.Logger
}

func (l *liquid) Emit(ops ...Activity) error {
	for _, o := range ops {
		go func(op Activity) {
			l.loop <- op
		}(o)
	}
	return nil
}

func (l *liquid) doActivity(op Activity) error {
	defer func() {
		if err := recover(); err != nil {
			l.logger.ErrorContext(l.c, "err", "error", err)
			debug.PrintStack()
		}
	}()

	if err := op.Do(l.c); err != nil {
		return err
	}

	// 发送事件
	go func() {
		l.ech <- op
	}()

	if err := op.Emit(l.c, l); err != nil {
		return err
	}

	return nil
}

func (l *liquid) doEventHander(op Activity) {
	defer func() {
		if err := recover(); err != nil {
			l.logger.ErrorContext(l.c, "err", "error", err)
			debug.PrintStack()
		}
	}()

	l.eh.Handle(l.c, op)
}

func (l *liquid) run() {
	for {
		select {
		case op, ok := <-l.loop:
			if !ok {
				return
			}

			if err := l.doActivity(op); err != nil {
				log.Print(err)
				continue
			}

		case op, ok := <-l.ech:
			if !ok {
				return
			}

			l.doEventHander(op)

		case <-l.c.Done():
			return
		}
	}
}

func (l *liquid) Run() error {
	for range l.size + 1 {
		go l.run()
	}
	return nil
}

func (l *liquid) Close() error {
	return nil
}
