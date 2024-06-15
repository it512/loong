package mongox

import "github.com/it512/loong"

type execData struct {
	ExecID string `json:"exec_id,omitempty"`

	ParentForkID string `json:"parent_fork_id,omitempty"` // 上一级fork , 顶级为InstID
	ForkID       string `json:"fork_id,omitempty"`        // 当前fork

	ForkTag string `json:"fork_tag,omitempty"` // 谁fork的
	JoinTag string `json:"join_tag,omitempty"` // 谁join的
	OutTag  string `json:"out_tag,omitempty"`  // fork的出口
	InTag   string `json:"in_tag,omitempty"`   // join的入口

	GwType   int `json:"gw_type,omitempty"`   // 网关类型 // 并行，包容
	ForkMode int `json:"fork_mode,omitempty"` // fork 模式

	Status int `json:"status,omitempty"`

	Input loong.Var `json:"input,omitempty"`
}
