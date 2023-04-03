package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
	"github.com/smart48ru/FaceIDApp/internal/domain"
)

// CreateEmployee implements openapi.ServerInterface
func (h *Handlers) CreateEmployee(c *gin.Context) {
	var r openapi.CreateEmployeeRequest

	err := c.BindJSON(&r)
	if err != nil {
		abortWithError(c, "error to parse employee request", err)
		return
	}
	employee := domain.Employee{
		Name:    r.Name,
		PhotoID: r.PhotoID,
		Meta:    r.Meta.AdditionalProperties,
	}

	id, err := h.staffApp.AddEmployee(c, employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.Error{Error: err.Error()})
		return
	}
	resp := openapi.CreateEmployeeResponse{
		ID:      id,
		Name:    r.Name,
		PhotoID: r.PhotoID,
		Meta:    r.Meta,
	}

	c.JSON(http.StatusOK, resp)
}
