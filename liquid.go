package loong

import (
	"log/slog"
)

type liquid struct {
	engine *Engine

	actCh chan Activity
	ech   chan Activity
	cmdCh chan Cmd

	size uint

	logger *slog.Logger
}

func (l *liquid) Background(cmds ...Cmd) error {
	for _, cmd := range cmds {
		go func(cmd Cmd) {
			l.cmdCh <- cmd
		}(cmd)
	}
	return nil
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
	err := l.engine.Txer.DoTrans(l.engine.ctx, func(tx TxContext) error {
		defer func() {
			if err := recover(); err != nil {
				if err = tx.Abort(tx); err != nil {
					l.logger.ErrorContext(l.engine.ctx, "activity panic", "error", err)
				}
			}
		}()

		var err error
		if err = op.Do(tx); err != nil {
			return err
		}

		l.sendEventAsync(op)

		if err = op.Emit(tx, l); err != nil {
			return err
		}

		if err = tx.Commit(tx); err != nil {
			return err
		}
		return nil
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

func (l *liquid) doActivity(op Activity) {
	defer func() {
		if err := recover(); err != nil {
			l.logger.ErrorContext(l.engine.ctx, "activity panic", "error", err)
		}
	}()

	var err error

	if err = op.Do(l.engine.ctx); err != nil {
		l.logger.ErrorContext(l.engine.ctx, "activity do error", "error", err)
	}

	l.sendEventAsync(op)

	if err = op.Emit(l.engine.ctx, l); err != nil {
		l.logger.ErrorContext(l.engine.ctx, "activity emit error", "error", err)
	}
}

func (l *liquid) doCmd(cmd Cmd) {
	defer func() {
		if err := recover(); err != nil {
			l.logger.ErrorContext(l.engine.ctx, "cmd panic", "error", err)
		}
	}()

	if err := cmd.Do(l.engine.ctx); err != nil {
		l.logger.ErrorContext(l.engine.ctx, "cmd error", "error", err)
	}
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
			l.doActivity(op)
		case op := <-l.ech:
			l.doEventHander(op)
		case cmd := <-l.cmdCh:
			l.doCmd(cmd)
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
