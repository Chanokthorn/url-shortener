package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"url-shortener/app/infrastructure/connector"
)

func main() {
	e := echo.New()
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDBName := os.Getenv("PG_DB_NAME")
	pgMigratePath := "file://./migrate/"
	pg := connector.NewPGConnector(pgHost, pgPort, pgUser, pgPassword, pgDBName)
	fmt.Print(pg)
	migrateDB(pgHost, pgPort, pgUser, pgPassword, pgDBName, pgMigratePath)
	e.Logger.Fatal(e.Start(":3000"))
}
