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
	v := &vote{Evaluator: eval}
	for _, u := range ut {
		v.numberOfInstances++

		if u.Status == STATUS_START { // 未投票的
			v.numberOfActiveInstances++
		}

		if u.Result != 0 { // 投票不通过
			v.numberOfTerminatedInstances++
		}
	}

	// 投票通过的 = 已投票 - 投票不通过的 = 总数 - 未投票的 - 投票不通过的
	v.numberOfCompletedInstances = v.numberOfInstances - v.numberOfActiveInstances - v.numberOfTerminatedInstances

	return v
}

func (v vote) Test(ctx context.Context, el string) (pass bool, err error) {
	if pass, _, err = eval2[bool](ctx, v, el, v.ToEnv()); err == nil { // 无错误
		if !pass && v.numberOfActiveInstances == 0 {
			panic(fmt.Errorf("投票已经结束，未能达成通过条件 %s", el))
		}
	}
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
