package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lyquocnam/zalora_icecream/controllers"
	"github.com/lyquocnam/zalora_icecream/lib"
	"net/http"
)

func loadRoutes(router *echo.Echo) {
	router.GET("/", controllers.GetHomeHandler)
	v1 := router.Group("/api/v1")

	v1.GET("/products", controllers.ProductGetListHandler)
	v1.GET("/products/:id", controllers.ProductGetByIDHandler)

	v1.POST("/login", controllers.LoginHandler)

	middleware.ErrJWTMissing = echo.NewHTTPError(http.StatusBadRequest, "you don't have permission to to this")

	authorized := v1.Group("", middleware.JWT([]byte(lib.Config.Secret)))
	authorized.POST("/products", controllers.ProductAddHandler)
	authorized.PUT("/products/:id", controllers.ProductUpdateHandler)
	authorized.DELETE("/products/:id", controllers.ProductDeleteHandler)
}
