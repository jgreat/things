package db

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/jgreat/things/util"
)

// Connect - Connect to database
func Connect() (*pg.DB, error) {
	url := util.GetEnv("DB_URL", "postgres://things:example@127.0.0.1:5432/things?sslmode=disable")
	options, err := pg.ParseURL(url)

	if err != nil {
		return nil, err
	}

	db := pg.Connect(options)

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to DB -- cool ")
	return db, nil
}
