package routes

import (
	"net/http"
	"tryout_echo/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type M map[string]interface{}

// custom error handling
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	data := M{
		"status" : code,
		"message" : err.(*echo.HTTPError).Message,
	}
	c.JSON(code,data)
}
//Main Route API
func Api() *echo.Echo {
	e := echo.New();
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Use(middleware.Logger())

	e.GET("/cek-api",func(c echo.Context) error {
		return c.JSON(http.StatusOK,"API ONLINE")
	})
	e.GET("/user",controllers.Index)

	return e
}