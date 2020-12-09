package models

import (
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// CreateTable -
func CreateTable(db *pg.DB, name string, model interface{}) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.Model(model).CreateTable(opts)
	if err != nil {
		return err
	}

	log.Printf("%s table created", name)
	return nil
}
