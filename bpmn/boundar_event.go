package bpmn

import "github.com/it512/loong/bpmn/zeebe"

type TBoundaryEvent struct {
	Id             string `xml:"id,attr"`
	CancelActivity bool   `xml:"cancelActivity,attr"`
	AttachedToRef  string `xml:"attachedToRef,attr"`

	OutgoingAssociation []string `xml:"outgoing"`

	MessageEventDefinition []TMessageEventDefinition `xml:"messageEventDefinition"`
	TimerEventDefinition   []TTimerEventDefinition   `xml:"timerEventDefinition"`

	Output []zeebe.TIoMapping `xml:"extensionElements>ioMapping>output"`
}
