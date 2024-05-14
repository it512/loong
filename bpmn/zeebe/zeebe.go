package zeebe

import (
	"encoding/xml"
)

type TAssignmentDefinition struct {
	XMLName         xml.Name
	Assignee        string `xml:"assignee,attr"`
	CandidateGroups string `xml:"candidateGroups,attr"`
	CandidateUsers  string `xml:"candidateUsers,attr"`
}

func (ad TAssignmentDefinition) GetCandidateGroups() []string {
	return splitTrim(ad.CandidateGroups, ",")
}

func (ad TAssignmentDefinition) GetCandidateUsers() []string {
	return splitTrim(ad.CandidateUsers, ",")
}

type TFormDefinition struct {
	XMLName xml.Name
	FormKey string `xml:"formKey,attr"`
	//FormID  string `xml:"formId,attr"`
}

type TIoMapping struct {
	XMLName xml.Name
	Source  string `xml:"source,attr"`
	Target  string `xml:"target,attr"`
}

type TProperty struct {
	XMLName xml.Name
	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr"`
}

type TTaskHeader struct {
	XMLName xml.Name
	Key     string `xml:"key,attr"`
	Value   string `xml:"value,attr"`
}

type TTaskDefinition struct {
	XMLName  xml.Name
	TypeName string `xml:"type,attr"`
	Retries  string `xml:"retries,attr"`
}

type LoopCharacteristics struct {
	XMLName          xml.Name
	InputCollection  string `xml:"inputCollection,attr"`
	InputElement     string `xml:"inputElement,attr"`
	OutputCollection string `xml:"outputCollection,attr"`
	OutputElement    string `xml:"outputElement,attr"`
}
