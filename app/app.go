package app

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lyquocnam/zalora_icecream/database"
	"github.com/lyquocnam/zalora_icecream/lib"
	"github.com/tylerb/graceful"
	"time"
)

// Start app
func Run() {
	lib.LoadConfig("config.yaml")
	lib.ConnectDatabase()
	database.Migrate()

	router := echo.New()
	router.Use(middleware.Gzip())
	router.Use(middleware.Recover())
	router.Use(middleware.RemoveTrailingSlash())

	// load all routes
	loadRoutes(router)

	// print app info
	printInfo()

	router.Server.Addr = fmt.Sprintf(`:%v`, lib.Config.AppPort)
	router.Logger.Fatal(graceful.ListenAndServe(router.Server, 5*time.Second))
}

func printInfo() {
	mode := "PRODUCTION"
	if lib.Config.Environment == "development" {
		mode = "DEVELOPMENT"
	} else if lib.Config.Environment == "staging" {
		mode = "STAGING"
	}

	fmt.Println("----------------------------------------")
	fmt.Println(fmt.Sprintf(`- %v`, lib.Config.AppName))
	fmt.Println(fmt.Sprintf(`- Version: %v`, lib.Config.AppVersion))
	fmt.Println(fmt.Sprintf(`- Environment: %v`, mode))
	fmt.Println(fmt.Sprintf(`- Author: %v`, lib.Config.Author))
	fmt.Println(fmt.Sprintf(`- Port: %v`, lib.Config.AppPort))
	fmt.Println("----------------------------------------")
}
