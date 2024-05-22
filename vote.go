package loong

type vote struct {
	numberOfInstances           int // The number of instances created.
	numberOfActiveInstances     int // The number of instances currently active.
	numberOfCompletedInstances  int // The number of instances already completed.
	numberOfTerminatedInstances int // The number of instances already terminated.
}

func (v *vote) Put(ut []UserTask) {
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
