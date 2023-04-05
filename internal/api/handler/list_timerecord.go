package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
	"net/http"
)

// ListTimerecord implements openapi.ServerInterface
func (h *Handlers) ListTimerecord(c *gin.Context) {
	timeRecords, err := h.timeRecordApp.ListTimeRecords(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.Error{Error: err.Error()})
		return
	}
	resp := make(openapi.ListTimerecordResponse, len(timeRecords))
	for i, timeRecord := range timeRecords {
		respTimeRecord := openapi.Timerecord{
			ID:       timeRecord.ID,
			Employee: timeRecord.Employee,
			EntryTime: openapi.TimerecordTime{
				Time:    timeRecord.EntryTime.Time,
				PhotoId: timeRecord.EntryTime.PhotoID,
			},
			ExitTime: &openapi.TimerecordTime{
				Time:    timeRecord.ExitTime.Time,
				PhotoId: timeRecord.ExitTime.PhotoID,
			},
		}
		resp[i] = respTimeRecord
	}

	c.JSON(http.StatusOK, resp)
}
