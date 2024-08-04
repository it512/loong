package loong

import (
	"context"

	"github.com/it512/loong/bpmn"
)

type SystemError struct {
	error
}

type BizErr struct {
	ErrorCode string
}

func NewBizErr(errCode string) *BizErr {
	return &BizErr{ErrorCode: errCode}
}

func (e BizErr) Error() string {
	return e.ErrorCode
}

func w(e error, sid string, v Variable) error {
	if be, ok := e.(*BizErr); ok {
		return &actErrOp{
			Variable: v,
			SourceID: sid,
			BizErr:   be,
		}
	}
	return e
}

type boundaryErrorEventActivity interface {
	Activity
	Match() bool
}

type actErrOp struct {
	Variable
	bpmn.TBoundaryEvent

	SourceID string

	*BizErr

	UnimplementedActivity
}

func (a *actErrOp) Match() bool {
	te, ok := a.Template.FindErrorByCode(a.ErrorCode)
	if !ok {
		return false
	}

	for _, be := range a.Template.Definitions.Process.BoundaryEvent {
		if be.AttachedToRef == a.SourceID {
			if ed, ok := be.GetErrorEventDefinition2(); ok {
				if ed.ErrorRef == te.ErrorCode {
					a.TBoundaryEvent = be
					return true
				}
			}
		}
	}

	return false
}

func (a *actErrOp) Emit(ctx context.Context, emt Emitter) error {
	return emt.Emit(fromOuter(ctx, a.Variable, a))
}
