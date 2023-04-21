package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
	"github.com/smart48ru/FaceIDApp/internal/domain"
)

// AddTimerecord implements openapi.ServerInterface
func (h *Handlers) AddTimerecord(c *gin.Context) {
	var r openapi.AddTimerecordRequest

	err := c.BindJSON(&r)
	if err != nil {
		abortWithError(c, "error to parse time record request", err)
		return
	}
	reqTimeRecord := domain.TimeRecord{
		Employee: r.Employee,
		EntryTime: domain.TimerecordTime{
			Time:    r.EntryTime.Time,
			PhotoID: r.EntryTime.PhotoId,
		},
	}

	timeRecordID, err := h.timeRecordApp.AddTimeRecord(c, reqTimeRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.Error{Error: err.Error()})
		return
	}
	respTimeRecord := openapi.AddTimerecordResponse{
		ID:       timeRecordID,
		Employee: r.Employee,
		EntryTime: openapi.TimerecordTime{
			PhotoId: r.EntryTime.PhotoId,
			Time:    r.EntryTime.Time,
		},
	}

	c.JSON(http.StatusOK, respTimeRecord)
}
