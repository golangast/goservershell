package server

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golangast/contribute/assets"
	"github.com/golangast/goservershell/src/routes"

	"github.com/Masterminds/sprig/v3"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func add() string {
	return "ere"
}

func Server() {
	e := echo.New()
	files, err := getAllFilenames(&assets.Assets)
	if err != nil {
		fmt.Print(err)
	}

	renderer := &TemplateRenderer{
		templates: template.Must(template.New("t").Funcs(template.FuncMap{
			"adds": add,
		}).Funcs(sprig.FuncMap()).ParseFS(assets.Assets, files...)),
	}

	e.Renderer = renderer
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: getFileSystem(assets.Assets),
		HTML5:      true,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	routes.Routes(e)

	// Route
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "",
	}))
	e.Use(middleware.BodyLimit("3M"))
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(30)))
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":5002"))

}

func GetAllFilePathsInDirectory(dirpath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func ParseDirectory(dirpath string) (*template.Template, error) {
	paths, err := GetAllFilePathsInDirectory(dirpath)
	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

var err error

func getFileSystem(TmplMainGo embed.FS) http.FileSystem {

	log.Print("using embed mode")
	fsys, err := fs.Sub(TmplMainGo, "assets/templates")
	if err != nil {
		log.Print(err)
	}

	return http.FS(fsys)
}

// https://gist.github.com/clarkmcc/1fdab4472283bb68464d066d6b4169bc
func getAllFilenames(efs *embed.FS) (files []string, err error) {
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}
