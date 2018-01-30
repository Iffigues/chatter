package main

import(
	"net/http"
	"github.com/figlet"

)

func init(){
	au := Router.PathPrefix("/figlet").Subrouter()
	au.Path("/").Methods("GET","POST").HandlerFunc(figlets)
	au.Path("/oauth").Methods("POST").HandlerFunc(NewFiglets)
}


type VV struct{
	Gg figlet.Response
	W  []string
	Checkbox []string
	F []string
}



func (resp *VV)makeArg(){
	resp.W = figlet.InputInt
	resp.Checkbox = figlet.InputCheckBox
	resp.F = figlet.F
}


func figlets(w http.ResponseWriter,r *http.Request){
	var(
		resp VV
		fortune = new(figlet.Coco)
	)
	if r.Method == "POST"{
		r.ParseForm()

	}
	if r.Method == "GET"{
		fortune.Input = "make feel a good day"
	}
	gr := figlet.CallFiglet(fortune)
	resp.Gg = gr
	resp.makeArg()
	aa := donne{}
	aa.Ar = resp
	jointure(r,w,aa,"layout.html","figlet.html")
	
}

func NewFiglets(w http.ResponseWriter,r *http.Request){
	var Mal *figlet.Coco
	grap(r,&Mal)
	vv := figlet.CallFiglet(Mal)
	sendJson(vv.Res,w)
}
