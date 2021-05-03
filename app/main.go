package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
	"url-shortener/app/infrastructure/connector"
	"url-shortener/app/infrastructure/middleware"
	"url-shortener/app/infrastructure/router"
	"url-shortener/app/interface/controller"
	"url-shortener/app/interface/repository"
	"url-shortener/app/usecase/interactor"
)

func parseEnvAsStringArray(name string) []string {
	s := os.Getenv(name)
	return strings.Split(s, ",")
}

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
	urlRepository := repository.NewUrlRepository(pg.DB)
	blacklist := parseEnvAsStringArray("BLACKLIST")
	urlInteractor := interactor.NewUrlInteractor(urlRepository, blacklist)
	urlClientController := controller.NewURLClientController(urlInteractor)
	urlAdminController := controller.NewURLAdminController(urlInteractor)
	urlRedirectController := controller.NewURLRedirectController(urlInteractor)
	adminTokenMiddleware := middleware.NewAdminTokenMiddleware(os.Getenv("ADMIN_TOKEN"))
	router.CreateURLClientRouter(e.Group("/client"), urlClientController)
	router.CreateURLAdminRouter(e.Group("/admin"), urlAdminController, adminTokenMiddleware)
	router.CreateURLRedirectRouter(e.Group("/redirect"), urlRedirectController)
	e.Logger.Fatal(e.Start(":3000"))
}
