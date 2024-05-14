package bpmn

import "github.com/it512/loong/bpmn/zeebe"

type TStartEvent struct {
	Id                     string                    `xml:"id,attr"`
	Name                   string                    `xml:"name,attr"`
	IsInterrupting         bool                      `xml:"isInterrupting,attr"`
	ParallelMultiple       bool                      `xml:"parallelMultiple,attr"`
	IncomingAssociation    []string                  `xml:"incoming"`
	OutgoingAssociation    []string                  `xml:"outgoing"`
	MessageEventDefinition []TMessageEventDefinition `xml:"messageEventDefinition"`
	TimerEventDefinition   []TTimerEventDefinition   `xml:"timerEventDefinition"`
	SignalEventDefinition  []TSignalEventDefinition  `xml:"signalEventDefinition"`
	Output                 []zeebe.TIoMapping        `xml:"extensionElements>ioMapping>output"`
	Properties             []zeebe.TProperty         `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>properties>property"`
}

func (startEvent TStartEvent) GetId() string {
	return startEvent.Id
}

func (startEvent TStartEvent) GetName() string {
	return startEvent.Name
}

func (startEvent TStartEvent) GetIncomingAssociation() []string {
	return startEvent.IncomingAssociation
}

func (startEvent TStartEvent) GetOutgoingAssociation() []string {
	return startEvent.OutgoingAssociation
}

func (startEvent TStartEvent) GetType() ElementType {
	return StartEvent
}

func (startEvent TStartEvent) HasMessageEventDefinition() bool {
	return len(startEvent.MessageEventDefinition) > 0
}

func (startEvent TStartEvent) GetErrorEventDefinition() TMessageEventDefinition {
	return startEvent.MessageEventDefinition[0]
}

func (startEvent TStartEvent) HasTimerEventDefinition() bool {
	return len(startEvent.TimerEventDefinition) > 0
}

func (startEvent TStartEvent) GetTimerEventDefinition() TTimerEventDefinition {
	return startEvent.TimerEventDefinition[0]
}

func (startEvent TStartEvent) HasSignalEventDefinition() bool {
	return len(startEvent.SignalEventDefinition) > 0
}

func (startEvent TStartEvent) GetSignalEventDefinition() TSignalEventDefinition {
	return startEvent.SignalEventDefinition[0]
}

func (startEvent TStartEvent) IsNormal() bool {
	return !(startEvent.HasMessageEventDefinition() ||
		startEvent.HasTimerEventDefinition() ||
		startEvent.HasSignalEventDefinition())
}
