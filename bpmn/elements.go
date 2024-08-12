package bpmn

type ElementType string

const (
	StartEvent             ElementType = "START_EVENT"
	EndEvent               ElementType = "END_EVENT"
	Task                   ElementType = "TASK"
	ServiceTask            ElementType = "SERVICE_TASK"
	BusinessRuleTask       ElementType = "BUSINESS_RULE_TASK"
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
