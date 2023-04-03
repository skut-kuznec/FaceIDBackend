package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
)

// DeleteEmployee implements openapi.ServerInterface
func (h *Handlers) DeleteEmployee(c *gin.Context, params openapi.DeleteEmployeeParams) {
	err := h.staffApp.DeleteEmployee(c, params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, openapi.Error{Error: err.Error()})
		return
	}
	resp := openapi.DeleteEmployeeResponse{
		Message: "success",
	}

	c.JSON(http.StatusOK, resp)
}
