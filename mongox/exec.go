package mongox

import (
	"context"

	"github.com/it512/loong"
)

// fork
func (m *Store) ForkExec(ctx context.Context, xs []loong.Exec) error {
	return nil
}

// join
func (m *Store) JoinExec(ctx context.Context, ex *loong.Exec) error {
	return nil
}

func (m *Store) LoadForks(ctx context.Context, forkID string) ([]loong.Exec, error) {
	return nil, nil
}

func (m *Store) LoadExec(ctx context.Context, execID string, ex *loong.Exec) error {
	return nil
}
