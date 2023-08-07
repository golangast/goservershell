package home

import (
	"net/http"

	"github.com/golangast/goservershell/internal/dbsql/gettable"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {

	user := gettable.Gettabledata("jim")

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"users": user,
	})

}
