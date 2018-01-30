package main

import(
	"github.com/calendar"
	"net/http"

)

type NY struct{
	Gg calendar.Response
	Ti []string
	Ch []string
}



func (resp *NY)makeArg(){
	resp.Ti = calendar.InputInt
	resp.Ch = calendar.InputCheckBox
}


func cal(w http.ResponseWriter,r *http.Request){
	var(
		resp NY
		calende = new(calendar.Coco)
	)
	if r.Method == "POST"{
		r.ParseForm()

	}
	if r.Method == "GET"{
	}
	gr := calendar.CallCal(calende)
	writeError(gr.Err)
	resp.Gg = gr
	resp.makeArg()
	aa := donne{}
	aa.Ar = resp
	jointure(r,w,aa,"layout.html","cal.html")
}

func NewCal(w http.ResponseWriter,r *http.Request){
	var Mal *calendar.Coco
	grap(r,&Mal)
	vv := calendar.CallCal(Mal)
	sendJson(vv.Res,w)
}

