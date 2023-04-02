package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
	"github.com/smart48ru/FaceIDApp/internal/app/imageapp"
	"github.com/smart48ru/FaceIDApp/internal/app/staffapp"
	"github.com/smart48ru/FaceIDApp/internal/app/timerecordapp"
)

var _ openapi.ServerInterface = &Handlers{}

type Handlers struct {
	imageApp      *imageapp.App
	staffApp      *staffapp.App
	timeRecordApp *timerecordapp.App
}

func New(
	imageApp *imageapp.App,
	staffApp *staffapp.App,
	timeRecordApp *timerecordapp.App,
) *Handlers {
	return &Handlers{
		imageApp:      imageApp,
		staffApp:      staffApp,
		timeRecordApp: timeRecordApp,
	}
}

func abortWithError(c *gin.Context, msgText string, err error) {
	log.Error().Err(err).Msg(msgText)

	ret := openapi.Error{
		Error: msgText,
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, ret)
}
