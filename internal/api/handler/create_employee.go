package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
)

// CreateEmployee implements openapi.ServerInterface
func (h *Handlers) CreateEmployee(c *gin.Context) {
	var r openapi.CreateEmployeeRequest

	err := c.BindJSON(&r)
	if err != nil {
		abortWithError(c, "error to parse empoyee request", err)
		return
	}

	// TODO: Вызвать логику, замапить ответ.
}
