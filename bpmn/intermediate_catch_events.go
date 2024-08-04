package bpmn

import "github.com/it512/loong/bpmn/zeebe"

type TIntermediateCatchEvent struct {
	Id                     string                    `xml:"id,attr"`
	Name                   string                    `xml:"name,attr"`
	IncomingAssociation    []string                  `xml:"incoming"`
	OutgoingAssociation    []string                  `xml:"outgoing"`
	MessageEventDefinition []TMessageEventDefinition `xml:"messageEventDefinition"`
	TimerEventDefinition   []TTimerEventDefinition   `xml:"timerEventDefinition"`
	LinkEventDefinition    []TLinkEventDefinition    `xml:"linkEventDefinition"`
	ParallelMultiple       bool                      `xml:"parallelMultiple"`
	Output                 []zeebe.TIoMapping        `xml:"extensionElements>ioMapping>output"`
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetId() string {
	return intermediateCatchEvent.Id
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetName() string {
	return intermediateCatchEvent.Name
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetIncomingAssociation() []string {
	return intermediateCatchEvent.IncomingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetOutgoingAssociation() []string {
	return intermediateCatchEvent.OutgoingAssociation
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetType() ElementType {
	return IntermediateCatchEvent
}

func (intermediateCatchEvent TIntermediateCatchEvent) HasLinkEventDefinition() bool {
	return len(intermediateCatchEvent.LinkEventDefinition) > 0
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetLinkEventDefinition() TLinkEventDefinition {
	return intermediateCatchEvent.LinkEventDefinition[0]
}

func (intermediateCatchEvent TIntermediateCatchEvent) GetLinkEventDefinition2() (link TLinkEventDefinition, ok bool) {
	if ok = intermediateCatchEvent.HasLinkEventDefinition(); ok {
		link = intermediateCatchEvent.GetLinkEventDefinition()
	}
	return
}
