package bpmn

import "github.com/it512/loong/bpmn/zeebe"

type TIntermediateThrowEvent struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`

	LinkEventDefinition    []TLinkEventDefinition    `xml:"linkEventDefinition"`
	SignalEventDefinition  []TSignalEventDefinition  `xml:"signalEventDefinition"`
	MessageEventDefinition []TMessageEventDefinition `xml:"messageEventDefinition"`

	TaskDefinition zeebe.TTaskDefinition `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>taskDefinition"`
	TaskHeaders    []zeebe.TTaskHeader   `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>taskHeaders>header"`
	Properties     []zeebe.TProperty     `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>properties>property"`

	Input  []zeebe.TIoMapping `xml:"extensionElements>ioMapping>input"`
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

func (intermediateThrowEvent TIntermediateThrowEvent) HasMessageEventDefinitio() bool {
	return len(intermediateThrowEvent.MessageEventDefinition) > 0
}

func (intermediateThrowEvent TIntermediateThrowEvent) GetLinkEventDefinition() TLinkEventDefinition {
	return intermediateThrowEvent.LinkEventDefinition[0]
}

func (intermediateThrowEvent TIntermediateThrowEvent) GetHasMessageEventDefinitio() TMessageEventDefinition {
	return intermediateThrowEvent.MessageEventDefinition[0]
}

func (serviceTask TIntermediateThrowEvent) GetIoInput() []zeebe.TIoMapping {
	return serviceTask.Input
}

func (serviceTask TIntermediateThrowEvent) GetIoOutput() []zeebe.TIoMapping {
	return serviceTask.Output
}

func (s TIntermediateThrowEvent) GetTaskHeader(key string) (string, bool) {
	return zeebe.GetTaskHeader(s.TaskHeaders, key)
}

func (s TIntermediateThrowEvent) GetProperty(name string) (string, bool) {
	return zeebe.GetProperty(s.Properties, name)
}
