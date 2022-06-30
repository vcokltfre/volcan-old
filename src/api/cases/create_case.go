package cases_api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vcokltfre/volcan/src/impl"
	"github.com/vcokltfre/volcan/src/utils"
)

func createCase(c echo.Context) error {
	var caseData createCaseData
	err := utils.DecodeAndValidate(c.Request().Body, &caseData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dbCase, err := impl.Interface.Cases.CreateCase(
		caseData.UserID,
		"",
		caseData.ModID,
		"",
		caseData.Type,
		caseData.Reason,
		caseData.MuteType,
		caseData.Metadata,
		caseData.Expires,
		caseData.Notified,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if !utils.Contains(caseTypes, dbCase.Type) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid case type")
	}

	return c.JSON(200, dbCase)
}
