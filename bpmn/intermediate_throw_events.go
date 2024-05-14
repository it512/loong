package bpmn

import "github.com/it512/loong/bpmn/zeebe"

type TIntermediateThrowEvent struct {
	Id                     string                   `xml:"id,attr"`
	Name                   string                   `xml:"name,attr"`
	IncomingAssociation    []string                 `xml:"incoming"`
	OutgoingAssociation    []string                 `xml:"outgoing"`
	LinkEventDefinition    []TLinkEventDefinition   `xml:"linkEventDefinition"`
	SignalEventDefinition  []TSignalEventDefinition `xml:"signalEventDefinition"`
	MessageEventDefinition []TMessageEventDefinition

	Output []zeebe.TIoMapping `xml:"extensionElements>ioMapping>output"`
}

func (intermediateThrowEvent TIntermediateThrowEvent) GetId() string {
	return intermediateThrowEvent.Id
}

func (intermediateThrowEvent TIntermediateThrowEvent) GetName() string {
	return intermediateThrowEvent.Name
}

func (intermediateThrowEvent TIntermediateThrowEvent) GetIncomingAssociation() []string {
	return intermediateThrowEvent.IncomingAssociation
}

func (intermediateThrowEvent TIntermediateThrowEvent) GetOutgoingAssociation() []string {
	return intermediateThrowEvent.OutgoingAssociation
}

func (intermediateThrowEvent TIntermediateThrowEvent) GetType() ElementType {
	return IntermediateThrowEvent
}

func (intermediateThrowEvent TIntermediateThrowEvent) HasLinkEventDefinition() bool {
	return len(intermediateThrowEvent.LinkEventDefinition) > 0
}

func (intermediateThrowEvent TIntermediateThrowEvent) GetLinkEventDefinition() TLinkEventDefinition {
	return intermediateThrowEvent.LinkEventDefinition[0]
}
