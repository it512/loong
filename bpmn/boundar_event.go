package bpmn

import "github.com/it512/loong/bpmn/zeebe"

type TBoundaryEvent struct {
	Id             string `xml:"id,attr"`
	Name           string `xml:"name,attr"`
	CancelActivity bool   `xml:"cancelActivity,attr"`
	AttachedToRef  string `xml:"attachedToRef,attr"`

	OutgoingAssociation []string `xml:"outgoing"`

	ErrorEventDefinition   []TErrorEventDefinition   `xml:"errorEventDefinition"`
	MessageEventDefinition []TMessageEventDefinition `xml:"messageEventDefinition"`
	TimerEventDefinition   []TTimerEventDefinition   `xml:"timerEventDefinition"`

	Output []zeebe.TIoMapping `xml:"extensionElements>ioMapping>output"`
}

func (e TBoundaryEvent) GetId() string {
	return e.Id
}

func (e TBoundaryEvent) GetName() string {
	return e.Name
}

func (e TBoundaryEvent) GetOutgoingAssociation() []string {
	return e.OutgoingAssociation
}
