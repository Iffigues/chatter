package main

import(
	"net/http"
	"github.com/fortuna"

)

func init(){
	auth := Router.PathPrefix("/fortune").Subrouter()
	auth.Path("/").Methods("GET","POST").HandlerFunc(fortunes)
	auth.Path("/oauth").Methods("POST").HandlerFunc(NewFortunes)
}

type TT struct{
	Fortune fortuna.Response
	Load []string
	M []string
	Input []string
}


func fortunes(w http.ResponseWriter,r *http.Request){
	var(
		resp TT 
		fortune = new(fortuna.TT)
	)
	if r.Method == "POST"{
		r.ParseForm()
	}
	if r.Method == "GET"{
		fortune.Arg = "e"
	}
	gr := fortuna.CallFortune(fortune)
	writeError(gr.Err)
	resp.Fortune = gr
	resp.Load = fortuna.OptionArray
	resp.Input = fortuna.RR
	resp.M = fortuna.InputString
	aa := donne{}
	aa.Ar = resp
	jointure(r,w,aa,"layout.html","fortune.html")
	
}
func NewFortunes(w http.ResponseWriter,r *http.Request){
	var (
		resp *fortuna.TT	
	)
	grap(r,&resp)
	gr := fortuna.CallFortune(resp)
	writeError(gr.Err)
	sendJson(gr.Res,w)
}

