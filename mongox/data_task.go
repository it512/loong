package mongox

import (
	"time"

	"github.com/it512/loong"
)

type userTaskData struct {
	TaskID string `json:"task_id,omitempty"`
	ProcID string `json:"proc_id,omitempty"`
	InstID string `json:"inst_id,omitempty"`
	ExecID string `json:"exec_id,omitempty"`

	BusiKey  string `json:"busi_key,omitempty"`
	BusiType string `json:"busi_type,omitempty"`

	FormKey string `json:"form_key,omitempty"`

	ActID   string `json:"act_id,omitempty"`
	ActName string `json:"act_name,omitempty"`

	Assignee        string   `json:"assignee,omitempty"`
	CandidateGroups []string `json:"candidate_groups,omitempty"`
	CandidateUsers  []string `json:"candidate_users,omitempty"`

	Operator string `json:"operator,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Manager  string `json:"manager,omitempty"`

	Input loong.Var `json:"input,omitempty"`

	Result int `json:"result,omitempty"`

	BatchNo string `json:"batch_no,omitempty"`

	Status int `json:"status,omitempty"`

	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`

	Version int `json:"version"`

	Tags loong.Var `json:"tags,omitempty"`
}

func usertaskdata_2_usertask_no_exec(ut userTaskData) loong.UserTask {
	return loong.UserTask{
		TaskID: ut.TaskID,

		FormKey: ut.FormKey,

		ActID:   ut.ActID,
		ActName: ut.ActName,

		Assignee:        ut.Assignee,
		CandidateGroups: ut.CandidateGroups,
		CandidateUsers:  ut.CandidateUsers,

		Operator: ut.Operator,
		Owner:    ut.Owner,
		Manager:  ut.Manager,

		Result: ut.Result,

		BatchNo: ut.BatchNo,

		Status: ut.Status,

		StartTime: ut.StartTime,
		EndTime:   ut.EndTime,

		Version: ut.Version,
	}
}

func usertask_2_usertaskdata(ut loong.UserTask) userTaskData {
	return userTaskData{
		TaskID: ut.TaskID,
		ProcID: ut.Exec.ProcInst.ProcID,
		InstID: ut.Exec.ProcInst.InstID,
		ExecID: ut.Exec.ExecID,

		BusiKey:  ut.Exec.ProcInst.BusiKey,
		BusiType: ut.Exec.ProcInst.BusiType,

		FormKey: ut.FormKey,

		ActID:   ut.ActID,
		ActName: ut.ActName,

		Assignee:        ut.Assignee,
		CandidateGroups: ut.CandidateGroups,
		CandidateUsers:  ut.CandidateUsers,

		Operator: ut.Operator,
		Owner:    ut.Owner,
		Manager:  ut.Manager,

		Input: ut.Variable.Input,

		Result: ut.Result,

		BatchNo: ut.BatchNo,

		Status: ut.Status,

		StartTime: ut.StartTime,
		EndTime:   ut.EndTime,

		Version: ut.Version,

		Tags: ut.Exec.ProcInst.Tags,
	}
}

func usertaskdata_ptr_2_usertask_ptr(ut *loong.UserTask, u *userTaskData) *loong.UserTask {
	ut.TaskID = u.TaskID
	ut.Exec.ProcInst.ProcID = u.ProcID
	ut.Exec.ProcInst.InstID = u.InstID
	ut.Exec.ExecID = u.ExecID

	ut.ProcInst.BusiKey = u.BusiKey
	ut.ProcInst.BusiType = u.BusiType

	ut.FormKey = u.FormKey

	ut.ActID = u.ActID
	ut.ActName = u.ActName

	ut.Assignee = u.Assignee
	ut.CandidateGroups = u.CandidateGroups
	ut.CandidateUsers = u.CandidateUsers

	ut.Operator = u.Operator
	ut.Owner = u.Owner
	ut.Manager = u.Manager

	ut.Variable.Input = u.Input

	ut.Result = u.Result

	ut.BatchNo = u.BatchNo

	ut.Status = u.Status

	ut.StartTime = u.StartTime
	ut.EndTime = u.EndTime

	ut.Version = u.Version

	ut.Exec.ProcInst.Tags = u.Tags

	return ut
}
