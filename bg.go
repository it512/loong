package loong

import "context"

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
	_ = b.engine.Txer.DoTx(b.engine.ctx, func(txCtx context.Context) error {
		return cmd.Do(txCtx)
	})
}
