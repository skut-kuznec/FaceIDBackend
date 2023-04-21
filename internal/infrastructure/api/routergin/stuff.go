package routergin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (g RouterGin) AddStaffEndPoint(context *gin.Context) {

}

// GetAllStaffEndPoint godoc
// @Summary health check
// @Schemes
// @Description health check
// @Tags Stuff Api
// @Accept json
// @Produce json
// @Success 200
// @Router /api/staff/all [get]
// @Success 200 {object} StuffAnswer "OK response"
func (g RouterGin) GetAllStaffEndPoint(c *gin.Context) {
	stuffAnswer := StuffAnswer{
		ID:      1,
		Name:    "Тест",
		PhotoID: 1,
	}
	c.JSON(http.StatusOK, stuffAnswer)
}

func (g RouterGin) UpdateStaffEndPoint(context *gin.Context) {

}

func (g RouterGin) DeleteStaffEndPoint(context *gin.Context) {

}

func (g RouterGin) GetStaffEndPoint(context *gin.Context) {

}

func (g RouterGin) RecognizeStaffEndPoint(context *gin.Context) {

}

func (g RouterGin) FindStaffEndPoint(context *gin.Context) {

}
