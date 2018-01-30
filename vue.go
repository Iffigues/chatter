package main

import(
	"html/template"
	"path/filepath"
	"net/http"

)
type Page struct {
	Title string
	Body  []byte
}

type vue struct{
	Title string
}

type donne struct{
	IsCo bool
	Name string
	Ar interface{}
}

func jointure(r *http.Request,w http.ResponseWriter,are donne,ar ...string){
	var joins []string
	are.IsCo = isRegister(r)
	for _,ok := range ar {
		joins = append(joins,filepath.Join("static", ok))

	}
	tmpl, _ := template.ParseFiles(joins...)
	
	tmpl.ExecuteTemplate(w, "layout", are)
}


func Title(r string)(v string){
	v = "gopiko "
	return
}
