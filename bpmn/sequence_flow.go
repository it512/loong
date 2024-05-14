package bpmn

import (
	"html"
	"strings"

	"github.com/it512/loong/bpmn/zeebe"
)

type TSequenceFlow struct {
	Id                  string        `xml:"id,attr"`
	Name                string        `xml:"name,attr"`
	SourceRef           string        `xml:"sourceRef,attr"`
	TargetRef           string        `xml:"targetRef,attr"`
	ConditionExpression []TExpression `xml:"conditionExpression"`

	Properties []zeebe.TProperty `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>properties>property"`
}

type TExpression struct {
	Text string `xml:",innerxml"`
	Type string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
}

func (e TSequenceFlow) HasConditionExpression() bool {
	return len(e.ConditionExpression) == 1 && len(strings.TrimSpace(e.GetConditionExpression())) > 0
}

func (e TSequenceFlow) GetConditionExpression() string {
	return html.UnescapeString(e.ConditionExpression[0].Text)
}

func (e TSequenceFlow) GetId() string {
	return e.Id
}
func (e TSequenceFlow) GetName() string {
	return e.Name
}
func (e TSequenceFlow) GetTargetRef() string {
	return e.TargetRef
}

func (e TSequenceFlow) GetSourceRef() string {
	return e.SourceRef
}

func (e TSequenceFlow) GetType() ElementType {
	return SequenceFlow
}
