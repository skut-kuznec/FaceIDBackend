package domain

import (
	"time"
)

type TimeRecord struct {
	Employee   int
	StartWork  time.Time
	FinishWork time.Time
}
