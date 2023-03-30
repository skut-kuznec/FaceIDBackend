package timerecord

import (
	"time"
)

// TimeRecord - структура записи
type TimeRecord struct {
	WorkerID   int       `json:"worker,omitempty"`
	StartWork  time.Time `json:"startWork,omitempty"`
	FinishWork time.Time `json:"finishWork,omitempty"`
}
