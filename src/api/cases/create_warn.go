package cases_api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vcokltfre/volcan/src/impl"
	"github.com/vcokltfre/volcan/src/utils"
)

func createWarn(c echo.Context) error {
	var warnData createWarnData
	err := utils.DecodeAndValidate(c.Request().Body, &warnData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dbWarn, err := impl.Interface.Cases.WarnUser(warnData.UserID, warnData.ModID, warnData.Reason, warnData.Notify, warnData.Metadata)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(200, dbWarn)
}
