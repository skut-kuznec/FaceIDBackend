package db

import (
	"FaceIDApp/internal/config"

	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewPostgres(cfg config.DB) (*sqlx.DB, error) {
	log.Debug().Msgf("Configuring Postgres")
	dbase, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d users=%s dbname=%s password=%s sslmode=%t",
		cfg.DBHost(), cfg.DBPort(), cfg.DBUser(), cfg.DBName(), cfg.DBPassword(), cfg.DBSSLMode()))
	if err != nil {
		return nil, err
	}
	err = dbase.Ping()
	if err != nil {
		return nil, err
	}
	// TODO уточнить структуру хранения данных описать создание таблици
	// _, err = dbase.Exec(`CREATE TABLE IF NOT EXISTS users (
	//	id int NOT NULL,
	//	created_at timestamptz NOT NULL,
	//	updated_at timestamptz NOT NULL,
	//	deleted_at timestamptz NULL,
	//	name varchar NOT NULL,
	//	"data" varchar NULL,
	//	perms int2 NULL,
	//	CONSTRAINT users_pk PRIMARY KEY (id)
	//)`)
	if err != nil {
		dbase.Close()
		return nil, err
	}

	return dbase, nil
}

//
//func NewPostgresOld(cfg config.DB) (*Store, error) {
//	log.Debug().Msgf("Configuring Postgres")
//	// dbase, err := sqlx.Open("postgres", dsn)
//	dbase, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d users=%s dbname=%s password=%s sslmode=%t",
//		cfg.DBHost(), cfg.DBPort(), cfg.DBUser(), cfg.DBName(), cfg.DBPassword(), cfg.DBSSLMode()))
//	if err != nil {
//		return nil, err
//	}
//	err = dbase.Ping()
//	if err != nil {
//		return nil, err
//	}
//	// TODO уточнить структуру хранения данных описать создание таблици
//	// _, err = dbase.Exec(`CREATE TABLE IF NOT EXISTS users (
//	//	id int NOT NULL,
//	//	created_at timestamptz NOT NULL,
//	//	updated_at timestamptz NOT NULL,
//	//	deleted_at timestamptz NULL,
//	//	name varchar NOT NULL,
//	//	"data" varchar NULL,
//	//	perms int2 NULL,
//	//	CONSTRAINT users_pk PRIMARY KEY (id)
//	//)`)
//	if err != nil {
//		dbase.Close()
//		return nil, err
//	}
//	store := &Store{
//		db: dbase,
//	}
//	return store, nil
//}
