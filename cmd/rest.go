/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/golangast/gentil/utility/ff"
	"github.com/golangast/goservershell/internal/loggers"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// restCmd represents the rest command
var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "A brief description of your command",
	Long:  `go run . rest -t post -n dog -f name.string age.int`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rest called")
		logger := loggers.CreateLogger()

		types, _ := cmd.Flags().GetString("types")
		name, _ := cmd.Flags().GetString("name")
		fields, _ := cmd.Flags().GetString("fields")
		folder, _ := cmd.Flags().GetString("folder")
		folderdir := folder + "/"

		field := GetPropDatatype(fields)

		d := Data{Fields: field, Lowercasename: name, Uppercasename: cases.Title(language.Und, cases.NoLower).String(name)}

		str := `
		package ` + name + `

		type {{.Uppercasename}} struct {
			{{range slice .Fields 0}}
			{{$a := splitn " " 2 . }}
				{{$a._0}} {{$a._1}}  ` + "`" + `param:"{{$a._0}}" query:"{{$a._0}}" header:"{{$a._0}}" form:"{{$a._0}}" json:"{{$a._0}}" xml:"{{$a._0}}" ` + "`" + `
		 	{{end}}
		}
		
		
		`
		handlerfile, err := gentil.Filefolder("./src/handler/restful/"+types+"/"+folderdir+name, name+".go")
		if err != nil {
			logger.Error(
				"trying to create handler file",
				slog.String("error: ", err.Error()),
			)
		}

		err = Writetemplateslice(str, handlerfile, d)
		if err != nil {
			logger.Error(
				"trying to update router.html",
				slog.String("error: ", err.Error()),
			)
		}

	},
}

func init() {
	rootCmd.AddCommand(restCmd)
	restCmd.Flags().StringP("types", "t", "", "Set your types")
	restCmd.Flags().StringP("name", "n", "", "Set your name")
	restCmd.Flags().StringP("fields", "f", "", "Set your fields")
	restCmd.Flags().StringP("folder", "o", "", "Set your folder")

}
func GetPropDatatype(prop string) []string {
	var property []string
	var types []string
	var field []string
	var strright string
	s := strings.Split(prop, " ")

	for _, ss := range s {
		sss := strings.Replace(ss, "\"", "", -1)
		property = append(property, TrimDot(sss))
		strright = strings.Replace(TrimDotright(sss), ".", "", 1)
		types = append(types, strright)
	}

	for a, str1_word := range property {
		for b, str2_word := range types {
			if a == b {
				field = append(field, str1_word+" "+str2_word)
			}
		}
	}
	return field
}
func TrimDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
func TrimDotright(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx:]
	}
	return s
}

func Writetemplateslice(temp string, f *os.File, d Data) error {
	functionMap := sprig.TxtFuncMap()
	dbmb := template.Must(template.New("t").Funcs(functionMap).Parse(temp))
	err := dbmb.Execute(f, d)
	if err != nil {
		return err
	}
	return nil
}

type Data struct {
	Fields        []string
	Lowercasename string
	Uppercasename string
}
