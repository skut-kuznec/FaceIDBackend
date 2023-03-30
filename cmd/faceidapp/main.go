package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/smart48ru/FaceIDApp/docs/static"
	"github.com/smart48ru/FaceIDApp/docs/swagger"
	"github.com/smart48ru/FaceIDApp/internal/api/handler"
	"github.com/smart48ru/FaceIDApp/internal/api/openapi"
	"github.com/smart48ru/FaceIDApp/internal/api/server"
	"github.com/smart48ru/FaceIDApp/internal/config"
	"github.com/smart48ru/FaceIDApp/internal/repository/imagerepo"
	"github.com/smart48ru/FaceIDApp/internal/repository/staffrepo"
	"github.com/smart48ru/FaceIDApp/internal/repository/timerecordrepo"
	"github.com/smart48ru/FaceIDApp/internal/service/imageservice"
	"github.com/smart48ru/FaceIDApp/internal/service/staffservice"
	"github.com/smart48ru/FaceIDApp/internal/service/timerecordservice"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Msgf("%s Loading config", err)
	}

	// Initializing repositories.
	staffRepo := staffrepo.New()
	imageRepo := imagerepo.New()
	timeRecordRepo := timerecordrepo.New()

	// Initializing services.
	staffService := staffservice.New(staffRepo)
	imageService := imageservice.New(imageRepo)
	timeRecordService := timerecordservice.New(timeRecordRepo)

	hn := handler.New(imageService, staffService, timeRecordService)

	if cfg.APIRRelease() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router = openapi.RegisterHandlers(router, hn)

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

	log.Info().Msgf("Running server on http://0.0.0.0:%d", cfg.API.APIPort())
	serv := server.NewServer(cfg, router)

	select {
	case err := <-serv.Start():
		log.Fatal().Err(err).Msg("error to start server")
	case <-ctx.Done():
		log.Info().Msg("Stopping server")
		serv.Stop()
	}
}
