package loong

import (
	"context"
	"fmt"
)

type vote struct {
	numberOfInstances           int // The number of instances created.
	numberOfActiveInstances     int // The number of instances currently active.
	numberOfCompletedInstances  int // The number of instances already completed.
	numberOfTerminatedInstances int // The number of instances already terminated.

	Evaluator
}

func newVote(ut []UserTask, eval Evaluator) *vote {
	var (
		numberOfActiveInstances     int // The number of instances currently active.
		numberOfCompletedInstances  int // The number of instances already completed.
		numberOfTerminatedInstances int // The number of instances already terminated.
	)

	for _, u := range ut {
		if u.Status == STATUS_START {
			numberOfActiveInstances++
		} else {
			switch {
			case u.Result > 0:
				numberOfCompletedInstances++
			case u.Result < 0:
				numberOfTerminatedInstances++
			default:
				panic("未投票")
			}
		}
	}

	return &vote{
		Evaluator: eval,

		numberOfInstances:           len(ut),
		numberOfActiveInstances:     numberOfActiveInstances,
		numberOfCompletedInstances:  numberOfCompletedInstances,
		numberOfTerminatedInstances: numberOfTerminatedInstances,
	}
}

func (v vote) Test(ctx context.Context, el string) (pass bool, err error) {
	if pass, _, err = eval2[bool](ctx, v.Evaluator, el, v.ToEnv()); err == nil { // 无错误
		if !pass && v.numberOfActiveInstances == 0 {
			panic(fmt.Errorf("投票已经结束，未能达成通过条件 %s", el))
		}
	}
	return
}

func (v vote) Eval(ctx context.Context, el string) (a any, err error) {
	if el == "" {
		return
	}

	a, err = v.Evaluator.Eval(ctx, el, v.ToEnv())
	return
}

func (v vote) ToEnv() Var {
	r := NewVar()

	r.
		Put("numberOfInstances", v.numberOfInstances).
		Put("numberOfActiveInstances", v.numberOfActiveInstances).
		Put("numberOfCompletedInstances", v.numberOfCompletedInstances).
		Put("numberOfTerminatedInstances", v.numberOfTerminatedInstances)

		// 兼容activity
	r.
		Put("nrOfInstances", v.numberOfInstances).
		Put("nrOfActiveInstances", v.numberOfActiveInstances).
		Put("nrOfCompletedInstances", v.numberOfCompletedInstances).
		Put("nrOfTerminatedInstances", v.numberOfTerminatedInstances)

	return r
}
