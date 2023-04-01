package domain

import (
	"time"
)

type TimeRecord struct {
	Employee   uint64
	StartWork  time.Time
	FinishWork time.Time
}
