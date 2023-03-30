package handler

import "github.com/gin-gonic/gin"

func (h *Handlers) RecognizeEmployee(c *gin.Context) {
	header, err := c.FormFile("file")
	_, _ = header, err
}
