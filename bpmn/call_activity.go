package bpmn

import "github.com/it512/loong/bpmn/zeebe"

type TCallActivy struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`

	CallElement zeebe.CalledElement `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>calledElement"`

	Input  []zeebe.TIoMapping `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>ioMapping>input"`
	Output []zeebe.TIoMapping `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>ioMapping>output"`
	// Input  []zeebe.TIoMapping `xml:"extensionElements>ioMapping>input"`
	// Output []zeebe.TIoMapping `xml:"extensionElements>ioMapping>output"`

	Properties []zeebe.TProperty `xml:"http://camunda.org/schema/zeebe/1.0 extensionElements>properties>property"`
}
