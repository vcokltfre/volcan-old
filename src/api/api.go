package api

import (
	"os"

	"github.com/labstack/echo/v4"
	cases_api "github.com/vcokltfre/volcan/src/api/cases"
)

func StartAPI() error {
	server := echo.New()
	server.HideBanner = true

	bind := ":8080"
	if env_bind := os.Getenv("BIND"); env_bind != "" {
		bind = env_bind
	}

	cases_api.Register(server)

	return server.Start(bind)
}
