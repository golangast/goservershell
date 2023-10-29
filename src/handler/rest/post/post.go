package post

import (
	"encoding/json"
	"net/http"

	"github.com/golangast/goservershell/internal/dbsql/user"
	"github.com/labstack/echo/v4"
)

func Posts(c echo.Context) error {
	// Bind the request body to a `User` struct.
	userr := new(user.Users)
	if err := c.Bind(userr); err != nil {
		return err
	}

	err := userr.Exists()
	if err != nil {
		return err
	} else {
		err = userr.Create()
		if err != nil {
			return err
		}
		// Marshal the `userr` struct to JSON.
		jsonData, err := json.Marshal(userr)
		if err != nil {
			return err
		}

		// Write the JSON data to the response.
		c.Response().Header().Set("Content-Type", "application/json")
		return c.JSON(http.StatusOK, jsonData)
	}

}
