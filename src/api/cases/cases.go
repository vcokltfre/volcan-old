package cases_api

import "github.com/labstack/echo/v4"

func Register(e *echo.Echo) {
	e.GET("/cases/:case_id", getCase)
	e.POST("/cases", createCase)
	e.POST("/cases/warn", createWarn)
}
