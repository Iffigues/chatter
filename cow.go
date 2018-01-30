package main

import(
	"net/http"	
	"github.com/cowsay"
)

func init(){
	auth := Router.PathPrefix("/cowsay").Subrouter()
	auth.Path("/").Methods("GET","POST").HandlerFunc(cow)
	auth.Path("/oauth").Methods("POST").HandlerFunc(newCow)
	aut := Router.PathPrefix("/toilet").Subrouter()
	aut.Path("/").Methods("GET","POST").HandlerFunc(toilets)
	aut.Path("/oauth").Methods("POST").HandlerFunc(NewToilet)
	
}

type CC struct{
	Gg cowsay.Response
	Load []string
	Type []string
	Style []string
	Animal []string
	W []string
	Sens  []string
}

func (resp *CC)makeArg(){
	resp.Style = cowsay.InputCheckBox
	resp.Load = cowsay.OptionArray
	resp.Animal = cowsay.Monstre
	resp.W = cowsay.InputInt
	resp.Sens = cowsay.InputString
	resp.Type = cowsay.CowChoice
}



func cow(w http.ResponseWriter,r *http.Request){
	var(
		resp CC
		fortune = new(cowsay.Coco)
	)
	if r.Method == "POST"{
		r.ParseForm()
		
	}
	if r.Method == "GET"{
		fortune.Input = "hello, je vais repeter ce que vous dites"
	}
	gr := cowsay.CallCow(fortune)
	resp.Gg = gr
	resp.makeArg()
	aa := donne{}
	aa.Ar = resp
	jointure(r,w,aa,"layout.html","cow.html")
	
}


func newCow(w http.ResponseWriter,r *http.Request){
	var Mal *cowsay.Coco
	grap(r,&Mal)
	vv := cowsay.CallCow(Mal)
	sendJson(vv.Res,w)
}


