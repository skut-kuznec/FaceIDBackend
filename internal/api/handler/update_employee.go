package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
	"github.com/smart48ru/FaceIDApp/internal/domain"
	"net/http"
)

// UpdateEmployee implements openapi.ServerInterface
func (h *Handlers) UpdateEmployee(c *gin.Context) {
	var r openapi.UpdateEmployeeRequest
	err := c.BindJSON(&r)
	if err != nil {
		abortWithError(c, "error to parse employee request", err)
		return
	}
	employee := domain.Employee{
		ID:      r.ID,
		Name:    r.Name,
		PhotoID: r.PhotoID,
		Meta:    r.Meta.AdditionalProperties,
	}

	updateEmployee, err := h.staffApp.UpdateEmployee(c, employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.Error{Error: err.Error()})
		return
	}
	resp := openapi.UpdateEmployeeResponse{
		ID:      updateEmployee.ID,
		Name:    updateEmployee.Name,
		PhotoID: updateEmployee.PhotoID,
		Meta: openapi.Meta{
			AdditionalProperties: updateEmployee.Meta,
		},
	}

	c.JSON(http.StatusOK, resp)
}
