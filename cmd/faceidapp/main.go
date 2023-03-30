package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/smart48ru/FaceIDApp/internal/config"
	"github.com/smart48ru/FaceIDApp/internal/infrastructure/api/handler"
	"github.com/smart48ru/FaceIDApp/internal/infrastructure/api/routergin"
	"github.com/smart48ru/FaceIDApp/internal/infrastructure/api/server"
	"github.com/smart48ru/FaceIDApp/internal/infrastructure/store/db"
	"github.com/smart48ru/FaceIDApp/internal/infrastructure/store/filestore"
	"github.com/smart48ru/FaceIDApp/internal/infrastructure/store/memstore"
	"github.com/smart48ru/FaceIDApp/internal/service"
	"github.com/smart48ru/FaceIDApp/internal/usecases/app/repos/imagerepo"
	"github.com/smart48ru/FaceIDApp/internal/usecases/app/repos/stuffrepo"
	"github.com/smart48ru/FaceIDApp/internal/usecases/app/repos/timerecordrepo"

	"github.com/rs/zerolog/log"
)

var uStore stuffrepo.StuffStore

var cStore timerecordrepo.TimeRecordStorage

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Msgf("%s Loading config", err)
	}
	// Create store
	switch cfg.DBEnable() {
	case true:
		dbCon, err := db.NewPostgres(cfg.DBConfig())
		if err != nil {
			log.Fatal().Msgf("not connect to DB | %s", err)
		}
		uStore = db.NewUSerStore(dbCon)
		cStore = db.NewTimeRecordStore(dbCon)
	case false:
		uStore = memstore.NewUseStore()
		cStore = memstore.NewCalendarStore()
	}

	fileStore := filestore.NewPhotoStore(cfg.FileUploadDir())
	// Create repo
	fileRepo := imagerepo.NewPhotoRepo(fileStore)
	calendarRepo := timerecordrepo.NewCalendarRepo(cStore)
	userRepo := stuffrepo.NewUsersRepo(uStore)
	// create service from management repository
	serviceApp := service.NewService(userRepo, calendarRepo, fileRepo)
	// Create handler
	userHandlers := handler.NewHandler(serviceApp)
	// Create router
	router := routergin.NewRouterGin(userHandlers, cfg.APIConfig())
	// Create & start http server
	serv := server.NewServer(cfg.APIConfig(), router)
	serv.Start()
	log.Info().Msg("Program Start")
	// wait graceful shutdown
	<-ctx.Done()
	log.Info().Msg("Program Stop")
	// stop http server
	serv.Stop()
	cancel()
}
