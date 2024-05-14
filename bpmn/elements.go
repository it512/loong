package bpmn

type ElementType string
type GatewayDirection string

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

	NilType ElementType = "Nil"

	Unspecified GatewayDirection = "Unspecified"
	Converging  GatewayDirection = "Converging"
	Diverging   GatewayDirection = "Diverging"
	Mixed       GatewayDirection = "Mixed"
)

type DefaultAttrElement interface {
	GetDefault() string
}

func (task TTask) GetId() string {
	return task.Id
}

func (task TTask) GetName() string {
	return task.Name
}

func (task TTask) GetIncomingAssociation() []string {
	return task.IncomingAssociation
}

func (task TTask) GetOutgoingAssociation() []string {
	return task.OutgoingAssociation
}

func (task TTask) GetType() ElementType {
	return Task
}

// -------------------------------------------------------------------------

func (parallelGateway TParallelGateway) GetId() string {
	return parallelGateway.Id
}

func (parallelGateway TParallelGateway) GetName() string {
	return parallelGateway.Name
}

func (parallelGateway TParallelGateway) GetIncomingAssociation() []string {
	return parallelGateway.IncomingAssociation
}

func (parallelGateway TParallelGateway) GetOutgoingAssociation() []string {
	return parallelGateway.OutgoingAssociation
}

func (parallelGateway TParallelGateway) GetType() ElementType {
	return ParallelGateway
}

func (parallelGateway TParallelGateway) IsParallel() bool {
	return true
}
func (parallelGateway TParallelGateway) IsExclusive() bool {
	return false
}

// -------------------------------------------------------------------------

func (exclusiveGateway TExclusiveGateway) GetId() string {
	return exclusiveGateway.Id
}

func (exclusiveGateway TExclusiveGateway) GetName() string {
	return exclusiveGateway.Name
}

func (exclusiveGateway TExclusiveGateway) GetIncomingAssociation() []string {
	return exclusiveGateway.IncomingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetOutgoingAssociation() []string {
	return exclusiveGateway.OutgoingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetType() ElementType {
	return ExclusiveGateway
}

func (exclusiveGateway TExclusiveGateway) GetDefault() string {
	return exclusiveGateway.Default
}

func (exclusiveGateway TExclusiveGateway) IsParallel() bool {
	return false
}
func (exclusiveGateway TExclusiveGateway) IsExclusive() bool {
	return true
}

// -------------------------------------------------------------------------

func (eventBasedGateway TEventBasedGateway) GetId() string {
	return eventBasedGateway.Id
}

func (eventBasedGateway TEventBasedGateway) GetName() string {
	return eventBasedGateway.Name
}

func (eventBasedGateway TEventBasedGateway) GetIncomingAssociation() []string {
	return eventBasedGateway.IncomingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetOutgoingAssociation() []string {
	return eventBasedGateway.OutgoingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetType() ElementType {
	return EventBasedGateway
}

func (eventBasedGateway TEventBasedGateway) IsParallel() bool {
	return false
}

func (eventBasedGateway TEventBasedGateway) IsExclusive() bool {
	return true
}
