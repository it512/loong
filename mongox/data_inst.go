package mongox

import (
	"time"

	"github.com/it512/loong"
)

type instData struct {
	InstID string `json:"inst_id,omitempty"`
	ProcID string `json:"proc_id,omitempty"`

	Starter string `json:"starter,omitempty"`

	BusiKey  string `json:"busi_key,omitempty"`
	BusiType string `json:"busi_type,omitempty"`

	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`

	Status int `json:"status,omitempty"`

	Init loong.Var `json:"init,omitempty"`

	Tags loong.Var `json:"tags,omitempty"`
}
