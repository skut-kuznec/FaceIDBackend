package domain

import "time"

type TimeRecord struct {
	ID        uint64
	Employee  uint64
	EntryTime struct {
		Time    time.Time
		PhotoId uint64
	}
	ExitTime time.Time
}
