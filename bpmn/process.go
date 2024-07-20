package bpmn

import "github.com/it512/loong/bpmn/zeebe"

type TProcess struct {
	Id                           string `xml:"id,attr"`
	Name                         string `xml:"name,attr"`
	ProcessType                  string `xml:"processType,attr"`
	IsClosed                     bool   `xml:"isClosed,attr"`
	IsExecutable                 bool   `xml:"isExecutable,attr"`
	DefinitionalCollaborationRef string `xml:"definitionalCollaborationRef,attr"`

	StartEvents            []TStartEvent             `xml:"startEvent"`
	EndEvents              []TEndEvent               `xml:"endEvent"`
	SequenceFlows          []TSequenceFlow           `xml:"sequenceFlow"`
	Tasks                  []TTask                   `xml:"task"`
	ServiceTasks           []TServiceTask            `xml:"serviceTask"`
	UserTasks              []TUserTask               `xml:"userTask"`
	ParallelGateway        []TParallelGateway        `xml:"parallelGateway"`
	ExclusiveGateway       []TExclusiveGateway       `xml:"exclusiveGateway"`
	IntermediateCatchEvent []TIntermediateCatchEvent `xml:"intermediateCatchEvent"`
	IntermediateTrowEvent  []TIntermediateThrowEvent `xml:"intermediateThrowEvent"`
	EventBasedGateway      []TEventBasedGateway      `xml:"eventBasedGateway"`

	// Properties []zeebe.TProperty `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>properties>property"`
	Properties []zeebe.TProperty `xml:"extensionElements>properties>property"`
}

func (s TProcess) GetProperty(name string) (string, bool) {
	return zeebe.GetProperty(s.Properties, name)
}
