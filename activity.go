package loong

import (
	"context"

	"github.com/it512/loong/bpmn"
)

const (
	STATUS_START  = 1
	STATUS_FINISH = 100
)

type ActivityType string

const (
	NotApplicable ActivityType = "N/A"
)

type BpmnElement interface{ bpmn.BaseElement }

type Activity interface {
	Do(context.Context) error
	Emit(context.Context, Emitter) error
	Type() ActivityType
}
