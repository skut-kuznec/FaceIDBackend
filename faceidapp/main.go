package main

import (
	"FaceIDApp/internal/config"
	"FaceIDApp/internal/infrastructure/api/handler"
	"FaceIDApp/internal/infrastructure/api/routergin"
	"FaceIDApp/internal/infrastructure/api/server"
	"FaceIDApp/internal/infrastructure/store/db"
	"FaceIDApp/internal/infrastructure/store/memstore"
	"FaceIDApp/internal/usecases/repos/usersrepo"
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("%s Loading config", err)
	}
	fmt.Printf("%+v", cfg)
	switch cfg.DBEnable() {
	case true:
		store, err := db.NewPostgres(cfg.DBConfig())
		if err != nil {
			log.Fatal().Msgf("not connect to DB | %s", err)
		}
		repo := usersrepo.NewUsersRepo(store)
		h := handler.NewHandler(repo)
		router := routergin.NewRouterGin(h, cfg.APIConfig())
		serv := server.NewServer(cfg.APIConfig(), router)
		serv.Start(repo)

		fmt.Println("Program Start")
		<-ctx.Done()
		fmt.Println("Program Stop")
		serv.Stop()
		cancel()

	case false:
		store := memstore.NewMemStore()
		repo := usersrepo.NewUsersRepo(store)
		h := handler.NewHandler(repo)
		router := routergin.NewRouterGin(h, cfg.APIConfig())
		serv := server.NewServer(cfg.APIConfig(), router)
		serv.Start(repo)

		fmt.Println("Program Start")
		<-ctx.Done()
		fmt.Println("Program Stop")
		serv.Stop()
		cancel()

	}

}
