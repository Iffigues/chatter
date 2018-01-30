package main

import (
	"net/http"
	"fmt"
	//"github.com/GeertJohan/go.rice"
)

func init(){
	auth := Router.PathPrefix("/").Subrouter()
	
	auth.Path("/").Methods("GET").HandlerFunc(home)
	auth.Path("/logout").Methods("GET").HandlerFunc(DecoClient)
	auth.Path("/login").Methods("GET","POST").HandlerFunc(CoClient)
	auth.Path("/signup").Methods("GET","POST").HandlerFunc(checkClient)
	au := Router.PathPrefix("/figlet").Subrouter()
	au.Path("/").Methods("GET","POST").HandlerFunc(figlets)
	au.Path("/oauth").Methods("POST").HandlerFunc(NewFiglets)
	a := Router.PathPrefix("/room").Subrouter()
	a.Path("/create").Methods("POST").HandlerFunc(createRoom)
	a.Path("/search").Methods("POST").HandlerFunc(searchRoom)
	a.Path("/connect").Methods("POST").HandlerFunc(connectRoom)
	authe := Router.PathPrefix("/rig").Subrouter()
	authe.Path("/").Methods("GET","POST").HandlerFunc(rigs)
	authe.Path("/oauth").Methods("POST").HandlerFunc(NewRig)
	authes := Router.PathPrefix("/calendar").Subrouter()
	authes.Path("/").Methods("GET","POST").HandlerFunc(cal)
	authes.Path("/oauth").Methods("POST").HandlerFunc(NewCal)
}



func errorHome(w http.ResponseWriter,r *http.Request){
	jointure(r,w,donne{},"layout.html","404.html")
}

func home(w http.ResponseWriter , r *http.Request){
	jointure(r,w,donne{},"layout.html","home.html")
}

func DecoClient(w http.ResponseWriter,r *http.Request){
	if deconnection(w,r){
		jointure(r,w,donne{},"layout.html","home.html")
		return
	}
	jointure(r,w,donne{},"layout.html","home.html")
}
func CoClient(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	if len(r.Form) == 0 {
		jointure(r,w,donne{},"layout.html","con.html")
		return
	}
	user := r.Form["login"][0]
	pass := r.Form["password"][0]
	if connection(w,r,user,pass){
		jointure(r,w,donne{},"layout.html","home.html")
	}
	jointure(r,w,donne{},"layout.html","con.html")
}

func checkClient(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	if len(r.Form) == 0 {
		jointure(r,w,donne{},"layout.html","signup.html")
		return
	}
	user := r.Form["login"][0]
	pass := r.Form["password"][0]
	email := r.Form["mail"][0]
	fmt.Println(r.Form)
	if inscription(w,r,user,pass,email){
		jointure(r,w,donne{},"layout.html","home.html")
	}else{
	jointure(r,w,donne{},"layout.html","signup.html")
	}
}


