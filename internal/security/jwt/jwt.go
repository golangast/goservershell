package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golangast/goservershell/internal/security/crypt"
	"github.com/labstack/echo/v4"
)

func CreateJWT(sessionname, sessionkey string) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = sessionname
	claims["authorized"] = true
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(sessionkey))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return t, nil
}
func AuthoCreate(c echo.Context, sessionname, sessionkey string) (string, string, error) {
	t, err := CreateJWT(sessionname, sessionkey)
	if err != nil {
		return "", "", err
	}
	cookie := new(http.Cookie)
	cookie.Name = sessionname
	cookie.Value = sessionkey
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return t, crypt.CreateHash(t), err
}
