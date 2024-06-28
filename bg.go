package loong

type Backgrounder interface {
	Background(...Cmd) error
}

type bg struct {
	engine *Engine
}

func (b bg) Background(cmds ...Cmd) error {
	for _, cmd := range cmds {
		go b.doCmd(cmd)
	}
	return nil
}

func (b bg) doCmd(cmd Cmd) {
	defer func() {
		if err := recover(); err != nil {
			b.engine.Logger.ErrorContext(b.engine.ctx, "err", "error", err)
		}
	}()

	if err := cmd.Do(b.engine.ctx); err != nil {
		b.engine.Logger.ErrorContext(b.engine.ctx, "cmd error", "error", err)
	}
}
