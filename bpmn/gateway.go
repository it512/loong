package bpmn

type TParallelGateway struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

func (parallelGateway TParallelGateway) GetId() string {
	return parallelGateway.Id
}

func (parallelGateway TParallelGateway) GetName() string {
	return parallelGateway.Name
}

func (parallelGateway TParallelGateway) GetIncomingAssociation() []string {
	return parallelGateway.IncomingAssociation
}

func (parallelGateway TParallelGateway) GetOutgoingAssociation() []string {
	return parallelGateway.OutgoingAssociation
}

func (parallelGateway TParallelGateway) GetType() ElementType {
	return ParallelGateway
}

func (parallelGateway TParallelGateway) IsParallel() bool {
	return true
}
func (parallelGateway TParallelGateway) IsExclusive() bool {
	return false
}

type TExclusiveGateway struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	Default             string   `xml:"default,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

func (exclusiveGateway TExclusiveGateway) GetId() string {
	return exclusiveGateway.Id
}

func (exclusiveGateway TExclusiveGateway) GetName() string {
	return exclusiveGateway.Name
}

func (exclusiveGateway TExclusiveGateway) GetIncomingAssociation() []string {
	return exclusiveGateway.IncomingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetOutgoingAssociation() []string {
	return exclusiveGateway.OutgoingAssociation
}

func (exclusiveGateway TExclusiveGateway) GetType() ElementType {
	return ExclusiveGateway
}

func (exclusiveGateway TExclusiveGateway) GetDefault() string {
	return exclusiveGateway.Default
}

func (exclusiveGateway TExclusiveGateway) IsParallel() bool {
	return false
}
func (exclusiveGateway TExclusiveGateway) IsExclusive() bool {
	return true
}

type TEventBasedGateway struct {
	Id                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
}

func (eventBasedGateway TEventBasedGateway) GetId() string {
	return eventBasedGateway.Id
}

func (eventBasedGateway TEventBasedGateway) GetName() string {
	return eventBasedGateway.Name
}

func (eventBasedGateway TEventBasedGateway) GetIncomingAssociation() []string {
	return eventBasedGateway.IncomingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetOutgoingAssociation() []string {
	return eventBasedGateway.OutgoingAssociation
}

func (eventBasedGateway TEventBasedGateway) GetType() ElementType {
	return EventBasedGateway
}

func (eventBasedGateway TEventBasedGateway) IsParallel() bool {
	return false
}

func (eventBasedGateway TEventBasedGateway) IsExclusive() bool {
	return true
}
