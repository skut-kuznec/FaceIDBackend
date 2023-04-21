package timerecordapp

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type TimeRecordRepo interface {
	Save(ctx context.Context, t domain.TimeRecord) (uint64, error)
	GetLastByEmpoyeeID(ctx context.Context, id uint64) (domain.TimeRecord, error)
	List(ctx context.Context) ([]domain.TimeRecord, error)
}

type App struct {
	repo TimeRecordRepo
}

func (a *App) AddTimeRecord(c *gin.Context, record domain.TimeRecord) (uint64, error) {
	id, err := a.repo.Save(c, record)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *App) ListTimeRecords(c *gin.Context) ([]domain.TimeRecord, error) {
	list, err := a.repo.List(c)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *App) GetLastByEmpoyeeID(c *gin.Context, id uint64) (domain.TimeRecord, error) {
	timeRecord, err := a.repo.GetLastByEmpoyeeID(c, id)
	if err != nil {
		return domain.TimeRecord{}, err
	}
	return timeRecord, nil
}

func New(repo TimeRecordRepo) *App {
	return &App{
		repo: repo,
	}
}
