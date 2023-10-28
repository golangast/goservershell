package resttemp

func Posted() string {
	var Post = `package {{.Lowercasename}}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
			type {{.Uppercasename}} struct {
				{{range slice .Fields 0}}
				{{$a := splitn " " 2 . }}
					{{$a._0}} {{$a._1}}  ` + "`" + `param:"{{$a._0}}" query:"{{$a._0}}" header:"{{$a._0}}" form:"{{$a._0}}" json:"{{$a._0}}" xml:"{{$a._0}}" ` + "`" + `
				 {{end}}
			}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	`
	return Post
}
