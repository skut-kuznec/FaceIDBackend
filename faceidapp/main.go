package main

import (
	"FaceIDApp/internal/config"
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
	if cfg.DBEnable() {
		//postgresDSN := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		//	cfg.DBUser(),
		//	cfg.DBPassword(),
		//	cfg.DBHost(),
		//	cfg.DBPort(),
		//	cfg.DBName(),
		//)
		store, err := db.NewPostgres(cfg.DBConfig())
		if err != nil {
			log.Fatal().Msgf("not connect to DB | %s", err)
		}
		repo := usersrepo.NewUsersRepo(store)
		fmt.Println(repo)
	} else {
		store := memstore.NewMemStore()
		repo := usersrepo.NewUsersRepo(store)
		fmt.Println(repo)
	}

	//TODO Уточнить структуру хэндлеров , создать и описать их

	fmt.Println("Program Start")
	<-ctx.Done()
	fmt.Println("Program Stop")
	cancel()
}
