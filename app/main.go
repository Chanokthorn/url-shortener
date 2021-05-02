package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"url-shortener/app/infrastructure/connector"
)

func main() {
	e := echo.New()
	pg := connector.NewPGConnector(
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB_NAME"),
	)
	fmt.Print(pg)
	e.Logger.Fatal(e.Start(":3000"))
}
