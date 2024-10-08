package bpmn

import (
	"github.com/it512/loong/bpmn/zeebe"
)

type TBusineeRuleTask struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	Default             string   `xml:"default,attr"`
	CompletionQuantity  int      `xml:"completionQuantity,attr"`
	IsForCompensation   bool     `xml:"isForCompensation,attr"`
	OperationRef        string   `xml:"operationRef,attr"`
	Implementation      string   `xml:"implementation,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`

	// Input          []zeebe.TIoMappingInput  `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>ioMapping>input"`
	// Output         []zeebe.TIoMappingOutput `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>ioMapping>output"`
	Input  []zeebe.TIoMapping `xml:"extensionElements>ioMapping>input"`
	Output []zeebe.TIoMapping `xml:"extensionElements>ioMapping>output"`

	TaskDefinition zeebe.TTaskDefinition `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>taskDefinition"`
	CalledDecision zeebe.CalledDecision  `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>calledDecision"`
	TaskHeaders    []zeebe.TTaskHeader   `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>taskHeaders>header"`
	Properties     []zeebe.TProperty     `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>properties>property"`
}

func (serviceTask TBusineeRuleTask) GetId() string {
	return serviceTask.Id
}

func (serviceTask TBusineeRuleTask) GetName() string {
	return serviceTask.Name
}

func (serviceTask TBusineeRuleTask) GetIncomingAssociation() []string {
	return serviceTask.IncomingAssociation
}

func (serviceTask TBusineeRuleTask) GetOutgoingAssociation() []string {
	return serviceTask.OutgoingAssociation
}

func (serviceTask TBusineeRuleTask) GetType() ElementType {
	return BusinessRuleTask
}

func (serviceTask TBusineeRuleTask) GetIoInput() []zeebe.TIoMapping {
	return serviceTask.Input
}

func (serviceTask TBusineeRuleTask) GetIoOutput() []zeebe.TIoMapping {
	return serviceTask.Output
}

func (s TBusineeRuleTask) GetTaskHeader(key string) (string, bool) {
	return zeebe.GetTaskHeader(s.TaskHeaders, key)
}

func (s TBusineeRuleTask) GetProperty(name string) (string, bool) {
	return zeebe.GetProperty(s.Properties, name)
}
