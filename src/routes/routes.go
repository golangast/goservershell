package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golangast/goservershell/internal/appdata"
	"github.com/golangast/goservershell/src/handler/get/home"
	"github.com/golangast/goservershell/src/handler/get/loginemail"
	"github.com/golangast/goservershell/src/handler/post/createuser"
	"github.com/golangast/goservershell/src/handler/restful/post"
	"github.com/labstack/echo/v4"
	// imports
)

type Stats struct {
	Exe  string
	Pid  string
	Ppid string
}

type Data struct {
	Proto   string
	Host    string
	Address string
	Method  string
	URL     string
	Path    string
	Routes  []string
	Stat    appdata.Stats
}

func Routes(e *echo.Echo) {

	e.GET("/home", home.Home)
	e.GET("/loginemail/:email/:sitetoken", loginemail.LoginEmail)
	//routes

	//post
	e.POST("/usercreate", createuser.Createuser)
	e.POST("/p", post.Posts)

	e.GET("/request", func(c echo.Context) error {
		req := c.Request()
		ss := e.Routes()
		var datas []string
		for key, _ := range ss {
			data, err := json.MarshalIndent(ss[key], "", "")
			if err != nil {
				return err
			}
			datas = append(datas, string(data))

		}
		st, err := appdata.GetAppData()
		if err != nil {
			fmt.Println(err)
		}

		d := Data{Proto: req.Proto, Host: req.Host, Address: req.RemoteAddr, Method: req.Method, Path: req.URL.Path, Routes: datas, Stat: st}
		//needed for nounce to be added to asset links for security to ensure they are those assets loading from this server
		nonce := c.Get("n")
		jsr := c.Get("jsr")
		cssr := c.Get("cssr")

		return c.Render(http.StatusOK, "request.html", map[string]interface{}{
			"data":  d,
			"nonce": nonce,
			"jsr":   jsr,
			"cssr":  cssr,
		})
	})
}
