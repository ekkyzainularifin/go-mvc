package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func Index(c echo.Context) error {
	data := M{
		"status" : http.StatusOK,
		"message": "Success",
		"data" : "index function",
	}
	return c.JSON(http.StatusOK,data)
}