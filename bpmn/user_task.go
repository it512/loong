package bpmn

import (
	"github.com/it512/loong/bpmn/zeebe"
)

// multiInstanceLoopCharacteristics
// xmins:activiti="http://activiti.org/bpmn"
type TMultiInstanceLoopCharacteristics struct {
	IsSequential bool `xml:"isSequential,attr"`

	Collection      string `xml:"http://activiti.org/bpmn collection,attr"`
	ElementVariable string `xml:"http://activiti.org/bpmn elementVariable,attr"`

	LoopCharacteristics zeebe.LoopCharacteristics `xml:"extensionElements>loopCharacteristics"`

	CompletionCondition TExpression `xml:"completionCondition"`
}

func (m TMultiInstanceLoopCharacteristics) GetCompletionCondition() string {
	return m.CompletionCondition.GetText()
}

func (m TMultiInstanceLoopCharacteristics) GetInputCollection() string {
	if m.LoopCharacteristics.InputCollection != "" {
		return m.LoopCharacteristics.InputCollection // for Camunda
	}

	return m.Collection // for Acviviti
}

func (m TMultiInstanceLoopCharacteristics) GetInputElement() string {
	if m.LoopCharacteristics.InputElement != "" {
		return m.LoopCharacteristics.InputElement // for Camunda
	}

	return m.ElementVariable // for Acviviti
}

type TUserTask struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`

	Input  []zeebe.TIoMapping `xml:"extensionElements>ioMapping>input"`
	Output []zeebe.TIoMapping `xml:"extensionElements>ioMapping>output"`

	AssignmentDefinition zeebe.TAssignmentDefinition `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>assignmentDefinition"`
	FormDefinition       zeebe.TFormDefinition       `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>formDefinition"`
	TaskHeaders          []zeebe.TTaskHeader         `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>taskHeaders>header"`
	Properties           []zeebe.TProperty           `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>properties>property"`

	MultiInstanceLoopCharacteristics []TMultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics"`
}

func (usertask TUserTask) HasMultiInstanceLoopCharacteristics() bool {
	return len(usertask.MultiInstanceLoopCharacteristics) > 0
}

func (usertask TUserTask) GetMultiInstanceLoopCharacteristics() TMultiInstanceLoopCharacteristics {
	if usertask.HasMultiInstanceLoopCharacteristics() {
		return usertask.MultiInstanceLoopCharacteristics[0]
	}
	panic("is not MultiInstance")
}

func (userTask TUserTask) GetId() string {
	return userTask.Id
}

func (userTask TUserTask) GetName() string {
	return userTask.Name
}

func (userTask TUserTask) GetIncomingAssociation() []string {
	return userTask.IncomingAssociation
}

func (userTask TUserTask) GetOutgoingAssociation() []string {
	return userTask.OutgoingAssociation
}

func (userTask TUserTask) GetType() ElementType {
	return UserTask
}

func (userTask TUserTask) GetInputMapping() []zeebe.TIoMapping {
	return userTask.Input
}

func (userTask TUserTask) GetOutputMapping() []zeebe.TIoMapping {
	return userTask.Output
}

func (userTask TUserTask) GetAssignmentAssignee() string {
	return userTask.AssignmentDefinition.Assignee
}

func (userTask TUserTask) GetAssignmentCandidateGroups() []string {
	return userTask.AssignmentDefinition.GetCandidateGroups()
}
func (u TUserTask) GetIoInput() []zeebe.TIoMapping {
	return u.Input
}

func (u TUserTask) GetIoOutput() []zeebe.TIoMapping {
	return u.Output
}

func (u TUserTask) GetTaskHeader(key string) (string, bool) {
	return zeebe.GetTaskHeader(u.TaskHeaders, key)
}

func (u TUserTask) GetProperty(name string) (string, bool) {
	return zeebe.GetProperty(u.Properties, name)
}
