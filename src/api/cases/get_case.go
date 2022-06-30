package cases_api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vcokltfre/volcan/src/impl"
)

func getCase(c echo.Context) error {
	caseID, err := strconv.Atoi(c.Param("case_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dbCase, err := impl.Interface.Cases.GetCase(caseID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(200, dbCase)
}
