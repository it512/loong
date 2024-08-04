package loong

import (
	"context"
	"log/slog"
)

type liquid struct {
	engine *Engine

	actCh chan Activity
	ech   chan Activity

	size uint

	logger *slog.Logger
}

func (l *liquid) Emit(ops ...Activity) error {
	for _, o := range ops {
		go func(op Activity) {
			l.actCh <- op
		}(o)
	}
	return nil
}

func (l *liquid) doActivityTx(op Activity) {
	err := l.engine.Txer.DoTx(l.engine.ctx, func(txCtx context.Context) (err error) {
		if err = op.Do(txCtx); err != nil {
			if e, ok := err.(boundaryErrorEventActivity); ok {
				if e.Match() {
					l.Emit(e)
				}
			}
			return
		}

		l.sendEventAsync(op)

		if err = op.Emit(txCtx, l); err != nil {
			return
		}
		return
	})

	if err != nil {
		l.logger.ErrorContext(l.engine.ctx, "activity panic", "error", err)
	}
}

func (l *liquid) sendEventAsync(op Activity) {
	// 发送事件
	go func() {
		l.ech <- op
	}()
}

func (l *liquid) doEventHander(op Activity) {
	defer func() {
		if err := recover(); err != nil {
			l.logger.ErrorContext(l.engine.ctx, "err", "error", err)
		}
	}()

	l.engine.Handle(l.engine.ctx, op)
}

func (l *liquid) loop() {
	for {
		select {
		case op := <-l.actCh:
			l.doActivityTx(op)
		case op := <-l.ech:
			l.doEventHander(op)
		case <-l.engine.ctx.Done():
			return
		}
	}
}

func (l *liquid) run() error {
	for range l.size + 1 {
		go l.loop()
	}
	return nil
}
