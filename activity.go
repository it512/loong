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

type UnimplementedActivity struct {
}

func (UnimplementedActivity) Do(ctx context.Context) error {
	return nil
}

func (UnimplementedActivity) Emit(ctx context.Context, emt Emitter) error {
	return nil
}

func (UnimplementedActivity) Type() ActivityType {
	return NotApplicable
}
