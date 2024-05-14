package loong

import (
	"context"
)

// forkMode
const (
	forkFork = 1
	newFork  = 10
	fullJoin = 66
	halfJoin = 88
)

// gw type
const (
	nogw      = 0
	exclusive = 1
	complexgw = 3
	eventbase = 5
	parallel  = 7
	inclusive = 9
)

type Forker interface {
	BpmnElement
	GetIncomingAssociation() []string
	GetOutgoingAssociation() []string
}

type gateway struct {
	Exec
	Forker
}

func (c *gateway) EmitExec(ctx context.Context, xs []Exec, emt Emitter) error {
	for _, ex := range xs {
		sf := fromExec(ex, ex.OutTag)
		emt.Emit(sf)
	}
	return nil
}

type exclusivGatewayOp struct {
	gateway
}

func (e exclusivGatewayOp) Do(ctx context.Context) error {
	switch e.Exec.GwType {
	case eventbase:
		e.Exec.GwType = nogw
	case nogw, exclusive:
		// skip
	case parallel, inclusive:
		panic("排他网关的前置网关不能为包容网关或并行网关")
	case complexgw:
		panic("目前不支持复杂网关")
	}

	return nil
}

func (e exclusivGatewayOp) Emit(ctx context.Context, emt Emitter) error {
	return e.EmitDefault(ctx, e.Forker, emt)
}
