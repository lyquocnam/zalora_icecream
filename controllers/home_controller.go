package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/lyquocnam/zalora_icecream/lib"
)

func GetHomeHandler(c echo.Context) error {
	return c.String(200, fmt.Sprintf(`%v v%v is working on /api`,
		lib.Config.AppName,
		lib.Config.AppVersion))
}
