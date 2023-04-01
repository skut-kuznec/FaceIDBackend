package router

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	router = openapi.RegisterHandlersWithOptions(router, si, openapi.GinServerOptions{
		Middlewares: []openapi.MiddlewareFunc{
			openapi.MiddlewareFunc(gin.Logger()),
		},
	})

	// Initializing swagger docs.
	router.StaticFS("/static", http.FS(static.SwaggerStatics))
	router.StaticFS("/docs", http.FS(swagger.UIPage))
	router.GET("/swagger.json", func(ctx *gin.Context) {
		swagger, err := openapi.GetSwagger()
		if err != nil {
			log.Error().Err(err).Msg("swagger error")
		}
		encoder := json.NewEncoder(ctx.Writer)
		_ = encoder.Encode(swagger)
	})

	upgrader := websocket.Upgrader{
		// Solve cross-domain problems
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// TODO: Получение сообщений из бизнес логики через канал. Ниже лишь пример.
	router.GET("/thirdparty/timerecordStream", func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer ws.Close()
		for {
			//read data from ws
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Error().Err(err).Msg("error to read message from websocket")
				break
			}
			log.Printf("recv: %s", message)

			//write ws data
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Error().Err(err).Msg("error to read message from websocket")
				break
			}
		}
	})

	return router
}
