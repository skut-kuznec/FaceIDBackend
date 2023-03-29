package main

import (
	"FaceIDApp/internal/config"
	"FaceIDApp/internal/infrastructure/api/handler"
	"FaceIDApp/internal/infrastructure/api/routergin"
	"FaceIDApp/internal/infrastructure/api/server"
	"FaceIDApp/internal/infrastructure/store/db"
	"FaceIDApp/internal/infrastructure/store/filestore"
	"FaceIDApp/internal/infrastructure/store/memstore"
	"FaceIDApp/internal/service"
	"FaceIDApp/internal/usecases/app/repos/calendarrepo"
	"FaceIDApp/internal/usecases/app/repos/photorepo"
	"FaceIDApp/internal/usecases/app/repos/usersrepo"

	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
)

var uStore usersrepo.UserStore

var cStore calendarrepo.CalendarStorage

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Msgf("%s Loading config", err)
	}
	// Проверяем работаем с базой или "memstorage"
	switch cfg.DBEnable() {
	case true:
		dbCon, err := db.NewPostgres(cfg.DBConfig())
		if err != nil {
			log.Fatal().Msgf("not connect to DB | %s", err)
		}
		uStore = db.NewUSerStore(dbCon)
		cStore = db.NewCalendarStore(dbCon)
	case false:
		uStore = memstore.NewUseStore()
		cStore = memstore.NewCalendarStore()
	}

	fileStore := filestore.NewPhotoStore("./upl")
	fileRepo := photorepo.NewPhotoRepo(fileStore)
	calendarRepo := calendarrepo.NewCalendarRepo(cStore)
	userRepo := usersrepo.NewUsersRepo(uStore)
	serviceApp := service.NewService(userRepo, calendarRepo, fileRepo)

	userHandlers := handler.NewHandler(serviceApp)
	router := routergin.NewRouterGin(userHandlers, cfg.APIConfig())
	serv := server.NewServer(cfg.APIConfig(), router)
	serv.Start(serviceApp)

	log.Info().Msg("Program Start")
	<-ctx.Done()
	log.Info().Msg("Program Stop")
	serv.Stop()
	cancel()
}
