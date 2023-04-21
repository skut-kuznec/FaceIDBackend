package domain

import "time"

type TimeRecord struct {
	ID        uint64
	Employee  uint64
	EntryTime TimerecordTime
	ExitTime  TimerecordTime
}

type TimerecordTime struct {
	Time    time.Time
	PhotoID uint64
}
