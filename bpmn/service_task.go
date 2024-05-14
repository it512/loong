package bpmn

import (
	"github.com/it512/loong/bpmn/zeebe"
)

type TServiceTask struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	Default             string   `xml:"default,attr"`
	CompletionQuantity  int      `xml:"completionQuantity,attr"`
	IsForCompensation   bool     `xml:"isForCompensation,attr"`
	OperationRef        string   `xml:"operationRef,attr"`
	Implementation      string   `xml:"implementation,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`

	//Input          []zeebe.TIoMappingInput  `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>ioMapping>input"`
	//Output         []zeebe.TIoMappingOutput `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>ioMapping>output"`
	Input  []zeebe.TIoMapping `xml:"extensionElements>ioMapping>input"`
	Output []zeebe.TIoMapping `xml:"extensionElements>ioMapping>output"`

	TaskDefinition zeebe.TTaskDefinition `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>taskDefinition"`
	TaskHeaders    []zeebe.TTaskHeader   `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>taskHeaders>header"`
	Properties     []zeebe.TProperty     `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>properties>property"`
}

func (serviceTask TServiceTask) GetId() string {
	return serviceTask.Id
}

func (serviceTask TServiceTask) GetName() string {
	return serviceTask.Name
}

func (serviceTask TServiceTask) GetIncomingAssociation() []string {
	return serviceTask.IncomingAssociation
}

func (serviceTask TServiceTask) GetOutgoingAssociation() []string {
	return serviceTask.OutgoingAssociation
}

func (serviceTask TServiceTask) GetType() ElementType {
	return ServiceTask
}
