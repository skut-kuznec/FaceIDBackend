package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
	"net/http"
)

// LastTimerecordByEmploeeID implements openapi.ServerInterface
func (h *Handlers) LastTimerecordByEmploeeID(c *gin.Context, params openapi.LastTimerecordByEmploeeIDParams) {
	last, err := h.timeRecordApp.GetLastByEmpoyeeID(c, params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.Error{Error: err.Error()})
		return
	}
	respTimeRecord := openapi.LastTimerecordByEmploeeIDResponse{
		ID:       last.ID,
		Employee: last.Employee,
		EntryTime: openapi.TimerecordTime{
			PhotoId: last.EntryTime.PhotoID,
			Time:    last.EntryTime.Time,
		},
		ExitTime: &openapi.TimerecordTime{
			PhotoId: last.ExitTime.PhotoID,
			Time:    last.ExitTime.Time,
		},
	}
	c.JSON(http.StatusOK, respTimeRecord)
}
