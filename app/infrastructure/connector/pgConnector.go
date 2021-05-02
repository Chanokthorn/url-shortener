package connector

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type PGConnector struct {
	DB *sqlx.DB
}

func NewPGConnector(host, port, user, password, dbName string) PGConnector {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	p := PGConnector{}
	p.DB = sqlx.NewDb(db, "postgres")
	err = p.DB.Ping()
	if err != nil {
		panic(err)
	}
	p.DB.SetConnMaxLifetime(0)
	p.DB.SetMaxIdleConns(0)
	log.Printf("%s, DB: %s", "Successfully connected to postgresql", dbName)
	return p
}
