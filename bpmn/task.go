package bpmn

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

func (task TTask) GetId() string {
	return task.Id
}

func (task TTask) GetName() string {
	return task.Name
}

func (task TTask) GetIncomingAssociation() []string {
	return task.IncomingAssociation
}

func (task TTask) GetOutgoingAssociation() []string {
	return task.OutgoingAssociation
}

func (task TTask) GetType() ElementType {
	return Task
}
