package main

import(
	"net/http"
	"github.com/toilet"

)

func init(){
	aut := Router.PathPrefix("/toilet").Subrouter()
	aut.Path("/").Methods("GET","POST").HandlerFunc(toilets)
	aut.Path("/oauth").Methods("POST").HandlerFunc(NewToilet)
}


type DD struct{
	Gg toilet.Response
	W  []string
	Type []string
	Checkbox []string
	F []string
}



func (resp *DD)makeArg(){
	resp.W = toilet.InputInt
	resp.Type = toilet.InputType
	resp.Checkbox = toilet.InputCheckBox
	resp.F = toilet.F
}


func toilets(w http.ResponseWriter,r *http.Request){
	var(
		resp DD
		fortune = new(toilet.Coco)
	)
	if r.Method == "POST"{
		r.ParseForm()

	}
	if r.Method == "GET"{
		fortune.Type = "gay"
		fortune.Input = "make feel a good day"
	}
	gr := toilet.CallCow(fortune)
	resp.Gg = gr
	resp.makeArg()
	aa := donne{}
	aa.Ar = resp
	jointure(r,w,aa,"layout.html","toilet.html")	
}

func NewToilet(w http.ResponseWriter,r *http.Request){
	var Mal *toilet.Coco
	grap(r,&Mal)
	vv := toilet.CallCow(Mal)
	sendJson(vv.Res,w)
}
