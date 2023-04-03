package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
)

// ListEmployees implements openapi.ServerInterface
func (h *Handlers) ListEmployees(c *gin.Context) {
	employees, err := h.staffApp.ListEmployee(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.Error{Error: err.Error()})
		return
	}
	resp := make(openapi.ListEmployeesResponse, len(employees))
	for i, employee := range employees {
		resp[i].ID = employee.ID
		resp[i].Name = employee.Name
		resp[i].PhotoID = employee.PhotoID
		resp[i].Meta.AdditionalProperties = employee.Meta
	}

	c.JSON(http.StatusOK, resp)
}
