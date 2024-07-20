package bpmn

import (
	"html"
)

// xmins:activiti="http://activiti.org/bpmn"

type TDefinitions struct {
	Id                 string     `xml:"id,attr"`
	Name               string     `xml:"name,attr"`
	TargetNamespace    string     `xml:"targetNamespace,attr"`
	ExpressionLanguage string     `xml:"expressionLanguage,attr"`
	TypeLanguage       string     `xml:"typeLanguage,attr"`
	Exporter           string     `xml:"exporter,attr"`
	ExporterVersion    string     `xml:"exporterVersion,attr"`
	Process            TProcess   `xml:"process"`
	Messages           []TMessage `xml:"message"`
	Errors             []TError   `xml:"error"`
	Signals            []TSignal  `xml:"signal"`
}

type TTask struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	Default             string   `xml:"default,attr"`
	CompletionQuantity  int      `xml:"completionQuantity,attr"`
	IsForCompensation   bool     `xml:"isForCompensation,attr"`
	OperationRef        string   `xml:"operationRef,attr"`
	Implementation      string   `xml:"implementation,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

type TParallelGateway struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

type TExclusiveGateway struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	Default             string   `xml:"default,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

type TEventBasedGateway struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

type TMessageEventDefinition struct {
	Id         string `xml:"id,attr"`
	MessageRef string `xml:"messageRef,attr"`
}

type TErrorEventDefinition struct {
	Id       string `xml:"id,attr"`
	ErrorRef string `xml:"errorRef,attr"`
}

type TSignalEventDefinition struct {
	Id        string `xml:"id,attr"`
	SignalRef string `xml:"signalRef,attr"`
}

type TTerminateEventDefinition struct {
	Id string `xml:"id,attr"`
}

type TTimerEventDefinition struct {
	Id           string        `xml:"id,attr"`
	TimeDuration TTimeDuration `xml:"timeDuration"`
}

type TLinkEventDefinition struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type TMessage struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type TSignal struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type TTimeDuration struct {
	XMLText string `xml:",innerxml"`
}

type TError struct {
	Id        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	ErrorCode string `xml:"errorCode,attr"`
}

func (e TError) GetId() string {
	return e.Id
}

type TExpression struct {
	Text string `xml:",innerxml"`
	Type string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
}

func (exp TExpression) GetText() string {
	return html.UnescapeString(exp.Text)
}
