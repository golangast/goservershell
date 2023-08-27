package welcome

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Welcome(c echo.Context) error {

	nonce := c.Get("n")

	return c.Render(http.StatusOK, "welcome.html", map[string]interface{}{
		"nonce": nonce,
	})

}
