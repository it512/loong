package bpmn

type DiagramError struct {
	Tag  string
	Code string
	error
}
