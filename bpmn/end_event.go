package bpmn

type TEndEvent struct {
	Id                       string                      `xml:"id,attr"`
	Name                     string                      `xml:"name,attr"`
	ErrorEventDefinition     []TErrorEventDefinition     `xml:"errorEventDefinition"`
	TerminateEventDefinition []TTerminateEventDefinition `xml:"terminateEventDefinition"`
	IncomingAssociation      []string                    `xml:"incoming"`
	OutgoingAssociation      []string                    `xml:"outgoing"`
}

func (endEvent TEndEvent) GetId() string {
	return endEvent.Id
}

func (endEvent TEndEvent) GetName() string {
	return endEvent.Name
}

func (endEvent TEndEvent) GetIncomingAssociation() []string {
	return endEvent.IncomingAssociation
}

func (endEvent TEndEvent) GetOutgoingAssociation() []string {
	return endEvent.OutgoingAssociation
}

func (endEvent TEndEvent) GetType() ElementType {
	return EndEvent
}

func (endEvent TEndEvent) HasErrorEventDefinition() bool {
	return len(endEvent.ErrorEventDefinition) > 0
}

func (endEvent TEndEvent) GetErrorEventDefinition() TErrorEventDefinition {
	return endEvent.ErrorEventDefinition[0]
}

func (endEvent TEndEvent) HasTerminateEventDefinition() bool {
	return len(endEvent.TerminateEventDefinition) > 0
}

func (endEvent TEndEvent) GetTerminateEventDefinition() TTerminateEventDefinition {
	return endEvent.TerminateEventDefinition[0]
}
