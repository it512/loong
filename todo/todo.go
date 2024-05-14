package todo

import (
	"context"
)

type Todo interface {
	QueryTaskPage(ctx context.Context, tp TodoQueryParam) (*TodoPageResult, error)
}

type TodoQueryParam struct {
	CandidateGroups []string `json:"candidate_groups,omitempty"`
	SudoGroups      []string `json:"sudo_groups,omitempty"`

	BusiType string `json:"busi_type,omitempty"`

	Last      string `json:"last,omitempty"`
	Size      int    `json:"size,omitempty"`
	Direction int    `json:"direction,omitempty"`
}

type Item struct {
	TaskID string `json:"task_id,omitempty"`

	FormKey string `json:"form_key,omitempty"`

	BusiKey  string `json:"busi_key,omitempty"`
	BusiType string `json:"busi_type,omitempty"`

	ActName string `json:"act_name,omitempty"`
	ActID   string `json:"act_id,omitempty"`

	CandidateGroups string `json:"candidate_groups,omitempty"`

	Owner string `json:"owner,omitempty"`

	Version int `json:"version,omitempty"`
	Status  int `json:"status,omitempty"`
}

type TodoPageResult struct {
	Items []*Item `json:"items,omitempty"`
	Max   string  `json:"max,omitempty"`
	Mix   string  `json:"min,omitempty"`
	Count int     `json:"count,omitempty"`
}
