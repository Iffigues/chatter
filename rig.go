package main

import(
	"net/http"
	"github.com/rig"
	
)


func init(){
	auth := Router.PathPrefix("/rig").Subrouter()
	auth.Path("/").Methods("GET","POST").HandlerFunc(rigs)
	auth.Path("/oauth").Methods("POST").HandlerFunc(NewRig)
}

type YY struct{
	Gg rig.Response
	Gender []string
	Nbr []string	
}

func (resp *YY)makeArg(){
	resp.Gender = rig.Genders
	resp.Nbr = rig.Nbrs
}

func rigs(w http.ResponseWriter,r *http.Request){
	var(
		resp YY
		fortune = new(rig.Coco)
	)
	if r.Method == "POST"{
		r.ParseForm()

	}
	if r.Method == "GET"{
	}
	gr := rig.CallRig(fortune)
	resp.Gg = gr
	resp.makeArg()
	aa := donne{}
	aa.Ar = resp
	jointure(r,w,aa,"layout.html","rig.html")
}


func NewRig(w http.ResponseWriter,r *http.Request){
	var Mal *rig.Coco
	grap(r,&Mal)
	vv := rig.CallRig(Mal)
	sendJson(vv.Res,w)
}
