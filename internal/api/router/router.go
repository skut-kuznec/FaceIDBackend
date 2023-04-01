package router

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/smart48ru/FaceIDApp/docs/static"
	"github.com/smart48ru/FaceIDApp/docs/swagger"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
)

func New(release bool, si openapi.ServerInterface) *gin.Engine {
	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router = openapi.RegisterHandlers(router, si)

	// Initializing swagger docs.
	router.Group("/static").StaticFS("", http.FS(static.SwaggerStatics))
	router.Group("/docs").StaticFS("", http.FS(swagger.UIPage))
	router.Group("").GET("/swagger.json", func(ctx *gin.Context) {
		swagger, err := openapi.GetSwagger()
		if err != nil {
			log.Error().Err(err).Msg("swagger error")
		}
		encoder := json.NewEncoder(ctx.Writer)
		_ = encoder.Encode(swagger)
	})

	return router
}
