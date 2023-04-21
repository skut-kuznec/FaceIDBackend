package main

import (
	"FaceIDApp/internal/api/handler"
	"FaceIDApp/internal/api/router"
	"FaceIDApp/internal/api/server"
	"FaceIDApp/internal/app/imageapp"
	"FaceIDApp/internal/app/staffapp"
	"FaceIDApp/internal/app/timerecordapp"
	"FaceIDApp/internal/config"
	"FaceIDApp/internal/repository/imagerepo"
	"FaceIDApp/internal/repository/staffrepo"
	"FaceIDApp/internal/repository/timerecordrepo"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Msgf("%s Loading config", err) //nolint: gocritic
	}

	// Initializing repositories.
	staffRepo := staffrepo.New()
	imageRepo := imagerepo.New()
	timeRecordRepo := timerecordrepo.New()

	// Initializing application logic.
	staffApp := staffapp.New(staffRepo)
	imageApp := imageapp.New(imageRepo)
	timeRecordApp := timerecordapp.New(timeRecordRepo)

	hn := handler.New(imageApp, staffApp, timeRecordApp)

	if cfg.APIRRelease() {
		gin.SetMode(gin.ReleaseMode)
	}

	rt := router.New(cfg.APIRRelease(), hn)

	log.Info().Msgf("Running server on http://0.0.0.0:%d", cfg.API.APIPort())
	serv := server.NewServer(cfg, rt)

	select {
	case err := <-serv.Start():
		log.Fatal().Err(err).Msg("error to start server")
	case <-ctx.Done():
		log.Info().Msg("Stopping server")
		serv.Stop()
	}
}
