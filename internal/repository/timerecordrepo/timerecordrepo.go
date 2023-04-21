package timerecordrepo

import (
	"context"
	"fmt"
	"github.com/smart48ru/FaceIDApp/internal/app/timerecordapp"
	"github.com/smart48ru/FaceIDApp/internal/domain"
	"sync"
)

var _ timerecordapp.TimeRecordRepo = &Repo{}

type Repo struct {
	mu   sync.Mutex
	m    map[uint64]domain.TimeRecord
	last map[uint64]uint64
	seq  uint64
}

func New() *Repo {
	return &Repo{
		m:    make(map[uint64]domain.TimeRecord),
		last: make(map[uint64]uint64),
	}
}

// GetLastByEmpoyeeID implements timerecordapp.TimeRecordRepo
func (r *Repo) GetLastByEmpoyeeID(ctx context.Context, id uint64) (domain.TimeRecord, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	select {
	case <-ctx.Done():
		return domain.TimeRecord{}, ctx.Err()
	default:
	}
	last, ok := r.last[id]
	if !ok {
		return domain.TimeRecord{}, fmt.Errorf("time record id=%d not exist", id)
	}
	lastEmployee, ok := r.m[last]
	if !ok {
		return domain.TimeRecord{}, fmt.Errorf("time record id=%d not exist", id)
	}
	return lastEmployee, nil
}

// List implements timerecordapp.TimeRecordRepo
func (r *Repo) List(ctx context.Context) ([]domain.TimeRecord, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	list := make([]domain.TimeRecord, len(r.m))
	for _, rec := range r.m {
		timeRecord := domain.TimeRecord{
			ID:       rec.ID,
			Employee: rec.Employee,
			EntryTime: domain.TimerecordTime{
				PhotoID: rec.EntryTime.PhotoID,
				Time:    rec.EntryTime.Time,
			},
			ExitTime: domain.TimerecordTime{
				PhotoID: rec.ExitTime.PhotoID,
				Time:    rec.ExitTime.Time,
			},
		}
		list = append(list, timeRecord)
	}
	return list, nil
}

// Save implements timerecordapp.TimeRecordRepo
func (r *Repo) Save(ctx context.Context, t domain.TimeRecord) (uint64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	r.seq += 1
	t.ID = r.seq
	r.m[r.seq] = t
	r.last[t.Employee] = r.seq

	return r.seq, nil
}
