package poster

import (
	"log/slog"

	"os"
	"text/template"

	"github.com/golangast/gentil/utility/ff"
	"github.com/golangast/goservershell/cmd/rest/resttemp"
	"github.com/golangast/goservershell/internal/loggers"

	"github.com/Masterminds/sprig/v3"
)

func Poster(d Fielddata, types, folderdir string) {
	logger := loggers.CreateLogger()

	handlerfile, err := gentil.Filefolder("./src/handler/restful/"+types+"/"+folderdir+d.Lowercasename, d.Lowercasename+".go")
	if err != nil {
		logger.Error(
			"trying to create handler file",
			slog.String("error: ", err.Error()),
		)
	}

	err = Writetemplateslice(resttemp.Posted(), handlerfile, d)
	if err != nil {
		logger.Error(
			"trying to update router.html",
			slog.String("error: ", err.Error()),
		)
	}

}

type Fielddata struct {
	Fields        []string
	CheckFields   []string
	Lowercasename string
	Uppercasename string
}

func Writetemplateslice(temp string, f *os.File, d Fielddata) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("t").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}
