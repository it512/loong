package loong

type SystemError struct {
	error
}

type BizErr struct {
	ErrorCode string
}

func (e BizErr) Error() string {
	return e.ErrorCode
}

/*
type actErrOp struct {
	Variable
	bpmn.TBoundaryEvent
	UnimplementedActivity

	error
}

func (s *actErrOp) Emit(ctx context.Context, emt Emitter) error {
	return emt.Emit(fromOuter(ctx, s.Exec, s))
}
*/
