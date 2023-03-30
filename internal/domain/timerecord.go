package domain

import (
	"time"
)

// TimeRecord - структура записи
type TimeRecord struct {
	WorkerID   int
	StartWork  time.Time
	FinishWork time.Time
}
