package resttemp

// types, _ := cmd.Flags().GetString("types")
// 		name, _ := cmd.Flags().GetString("name")
// 		fields, _ := cmd.Flags().GetString("fields")
// 		folder, _ := cmd.Flags().GetString("folder")
// 		folderdir := folder + "/"

// 		field := GetPropDatatype(fields)
// 		m := make(map[string][]string)
// 		m["field"] = field

// 		str := `
// 		package ` + name + `

// 		type ` + cases.Title(language.Und, cases.NoLower).String(name) + ` struct {
// 			{{range slice .field 0}}
// 			{{$a := splitn " " 2 . }}
// 				{{$a._0}} {{$a._1}}  ` + "`" + `param:"{{$a._0}}" query:"{{$a._0}}" header:"{{$a._0}}" form:"{{$a._0}}" json:"{{$a._0}}" xml:"{{$a._0}}" ` + "`" + `
// 		 	{{end}}
// 		}

// 		`
