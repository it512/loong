package bpmn

type ElementType string

func (e ElementType) Type() (string, error) {
	return string(e), nil
}

func (e ElementType) String() string {
	return string(e)
}

func (e ElementType) Is(ele ElementType) bool {
	return e == ele
}

func AsElementType(s string) ElementType {
	return ElementType(s)
}

const (
	StartEvent             ElementType = "START_EVENT"
	EndEvent               ElementType = "END_EVENT"
	Task                   ElementType = "TASK"
	ServiceTask            ElementType = "SERVICE_TASK"
	UserTask               ElementType = "USER_TASK"
	ParallelGateway        ElementType = "PARALLEL_GATEWAY"
	ExclusiveGateway       ElementType = "EXCLUSIVE_GATEWAY"
	InclusiveGateway       ElementType = "INCLUSIVE_GATEWAY"
	EventBasedGateway      ElementType = "EVENT_BASED_GATEWAY"
	IntermediateCatchEvent ElementType = "INTERMEDIATE_CATCH_EVENT"
	IntermediateThrowEvent ElementType = "INTERMEDIATE_THROW_EVENT"
	SequenceFlow           ElementType = "SEQUENCE_FLOW"
)

type DefaultAttrElement interface {
	GetDefault() string
}
