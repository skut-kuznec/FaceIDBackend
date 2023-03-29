package workcalendar

import (
	"time"

	"github.com/google/uuid"
)

// WorkCalendar - Календарь рабочих дней
type WorkCalendar struct {
	WorkDay    time.Time `json:"work_day,omitempty"`
	Worker     uuid.UUID `json:"worker,omitempty"`
	StartWork  time.Time `json:"startWork,omitempty"`
	InWork     bool      `json:"InWork,omitempty"`
	FinishWork time.Time `json:"finishWork,omitempty"`
}
