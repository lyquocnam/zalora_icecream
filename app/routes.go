package app

import (
	"github.com/labstack/echo"
	"github.com/lyquocnam/zalora_icecream/controllers"
)

func loadRoutes(router *echo.Echo) {
	router.GET("/", controllers.GetHomeHandler)
	v1 := router.Group("/api/v1")

	v1.GET("/products", controllers.ProductGetListHandler)
}
