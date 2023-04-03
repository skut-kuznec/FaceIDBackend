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
		respEmployee := openapi.Employee{
			ID:      employee.ID,
			Name:    employee.Name,
			PhotoID: employee.PhotoID,
			Meta: openapi.Meta{
				AdditionalProperties: employee.Meta,
			},
		}
		resp[i] = respEmployee
	}

	c.JSON(http.StatusOK, resp)
}
