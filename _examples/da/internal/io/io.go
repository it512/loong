package io

import (
	"context"

	"github.com/it512/loong"
)

type Io struct {
}

func (Io) Call(ctx context.Context, o loong.IoTasker) error {
	return loong.NewBizErr("error-01")
}
